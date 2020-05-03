package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/szeist/roomba-go/pkg/roomba"
)

func InteractiveMode(roombaClient *roomba.Roomba) {
	scanner := bufio.NewScanner(os.Stdin)
	name := roombaClient.GetStatus(2000).Name.(string)
	printHelp()
	printPrompt(name)
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "quit":
			return
		case "exit":
			return
		case "status":
			Status(roombaClient)
		case "help":
			printHelp()
		case "?":
			printHelp()
		default:
			SendCommand(roombaClient, text)
		}
		printPrompt(name)
	}
}

func printPrompt(name string) {
	fmt.Printf("%s > ", name)
}

func printHelp() {
	fmt.Println("status     - get roomba status")
	fmt.Println("exit, quit - exits application")
	fmt.Println("help, ?    - prints this help")
	fmt.Println("\ncommands:\n\t", strings.Join(roomba.GetSupportedCommands(), ", "))
}
