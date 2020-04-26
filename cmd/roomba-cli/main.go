package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/szeist/roomba-go/pkg/config"
	"github.com/szeist/roomba-go/pkg/roomba"
	"github.com/szeist/roomba-go/pkg/status"
)

var roombaClient *roomba.Roomba
var roombaStatus *status.Status

func main() {
	interactiveFlag := flag.Bool("interactive", false, "interactive mode")
	cmdFlag := flag.String("cmd", "", "roomba command")
	statusFlag := flag.Bool("status", false, "get roomba status")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		exit(1)
	}

	cfg := config.NewFromEnv("ROOMBA_")
	roombaClient = roomba.New(cfg)
	err := roombaClient.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		exit(2)
	}

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
