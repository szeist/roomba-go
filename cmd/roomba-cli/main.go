package main

import (
	"fmt"
	"os"

	"github.com/szeist/roomba-go/internal/cliapp"
	"github.com/szeist/roomba-go/internal/commands"
	"github.com/szeist/roomba-go/pkg/osargs"
)

func main() {
	app := &cliapp.App{}
	defer cleanup(app)

	checkArgs()

	subcommand := os.Args[1]

	switch subcommand {
	case "discover":
		commands.Discover()
	case "get-password":
		commands.GetPassword()
	case "status":
		initApp(app, subcommand)
		commands.Status(app.GetClient())
	case "cmd":
		initApp(app, subcommand)
		commands.SendCommand(app.GetClient(), os.Args[2])
	case "interactive":
		initApp(app, subcommand)
		commands.InteractiveMode(app.GetClient())
	default:
		printHelp()
		os.Exit(1)
	}
}

func initApp(app *cliapp.App, appName string) {
	app.ParseFlags(appName)

	err := app.Init()
	fail(err)

	err = osargs.MaskCmdlineArg("password")
	fail(err)
}

func checkArgs() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(1)
	}
}

func cleanup(app *cliapp.App) {
	err := app.Cleanup()
	fail(err)
}

func printHelp() {
	fmt.Fprintf(os.Stderr, `Usage: %s [command] [--help | arguments]
Available commands:
	discover      Discovers roomba on connected networks
	get-password  Get roomba password
	status        Get latest status info
	cmd           Send command to roomba
	interactive   Interactive mode
`, os.Args[0])
}

func fail(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
