package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/szeist/roomba-go/pkg/config"
	"github.com/szeist/roomba-go/pkg/discover"
	"github.com/szeist/roomba-go/pkg/roomba"
	"github.com/szeist/roomba-go/pkg/roombapass"
	"github.com/szeist/roomba-go/pkg/status"
)

var roombaClient *roomba.Roomba
var roombaStatus *status.Status

func main() {
	interactiveFlag := flag.Bool("interactive", false, "interactive mode")
	discoverFlag := flag.Bool("discover", false, "discover roombas on network")
	discoverTimeoutFlag := flag.Int("discover-timeout", 2, "roomba discovery timeout in seconds")
	getPasswordFlag := flag.Bool("get-password", false, "get roomba password")
	hostFlag := flag.String("host", "", "roomba ip address")
	cmdFlag := flag.String("cmd", "", "roomba command")
	statusFlag := flag.Bool("status", false, "get roomba status")
	debugFlag := flag.Bool("debug", false, "enable debug messages")
	captureStatusFlag := flag.String("capture-status", "", "capture status to file")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		exit(1)
	}

	if *discoverFlag {
		discoverCmd(*discoverTimeoutFlag)
		exit(0)
	}

	if *getPasswordFlag {
		if *hostFlag == "" {
			fmt.Fprintln(os.Stderr, "Specify the roomba ip address in the host flag")
			exit(1)
		}
		getPasswordCmd(*hostFlag)
		exit(0)
	}

	var statusWriter io.Writer
	if *captureStatusFlag != "" {
		file, err := os.Create(*captureStatusFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open file %s: %s", *captureStatusFlag, err.Error())
			exit(1)
		}
		statusWriter = file
	}

	initRoombaClient(*debugFlag, &statusWriter)

	if *interactiveFlag {
		interactiveMode()
	} else {
		if *cmdFlag != "" {
			sendCommand(*cmdFlag)
		}
		if *statusFlag {
			statusCmd()
		}
	}

	roombaClient.Disconnect()
}

func initRoombaClient(debug bool, statusWriter *io.Writer) {
	cfg := config.NewFromEnv("ROOMBA_")
	cfg.Debug = debug
	if statusWriter != nil {
		cfg.StatusWriter = statusWriter
	}
	roombaClient = roomba.New(cfg)
	err := roombaClient.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		exit(2)
	}
}

func interactiveMode() {
	scanner := bufio.NewScanner(os.Stdin)
	printHelp()
	printPrompt()
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "quit":
			exit(0)
		case "exit":
			exit(0)
		case "status":
			statusCmd()
		case "help":
			printHelp()
		case "?":
			printHelp()
		default:
			sendCommand(text)
		}
		printPrompt()
	}
}

func sendCommand(cmd string) {
	err := roombaClient.SendCommand(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	} else {
		fmt.Println("sent")
	}
}

func statusCmd() {
	status := roombaClient.GetStatus(10000)
	roombaStatus = &status
	jsonStatus, err := json.Marshal(status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
	fmt.Println(string(jsonStatus))
}

func discoverCmd(timeout int) {
	results, err := discover.Discover(time.Duration(timeout) * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
	if len(results) == 0 {
		fmt.Println("No Roomba found :(")
	} else {
		fmt.Println("Discovered Roombas:")
		for _, r := range results {
			fmt.Println(r)
		}
	}
}

func getPasswordCmd(host string) {
	fmt.Fprintln(os.Stderr, "Press the home button on the roomba for at least 2 seconds.\nPress enter when ready!")
	fmt.Scanln()
	password, err := roombapass.GetPassword(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		exit(1)
	}
	fmt.Println(password)
}

func exit(code int) {
	if roombaClient != nil && roombaClient.IsConnected() {
		roombaClient.Disconnect()
	}
	os.Exit(code)
}

func printPrompt() {
	status := roombaClient.GetStatus(250)
	roombaStatus = &status
	name := roombaStatus.Name
	if name == nil {
		name = ""
	}
	fmt.Printf("%s > ", name)
}

func printHelp() {
	fmt.Println("status     - get roomba status")
	fmt.Println("exit, quit - exits application")
	fmt.Println("help, ?    - prints this help")
	fmt.Println("\ncommands:\n\t", strings.Join(roomba.GetSupportedCommands(), ", "))
}
