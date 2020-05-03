package cliapp

import (
	"flag"
	"os"

	"github.com/szeist/roomba-go/pkg/config"
	"github.com/szeist/roomba-go/pkg/roomba"
)

const passwordArgName string = "password"

type App struct {
	host             string
	user             string
	password         string
	debug            bool
	stateLogFileName string
	stateLogFile     *os.File
	roombaClient     *roomba.Roomba
}

func (a *App) ParseFlags(appName string) {
	flagSet := flag.NewFlagSet(appName, flag.ExitOnError)
	flagSet.StringVar(&a.host, "host", "", "roomba IP address")
	flagSet.StringVar(&a.user, "user", "", "roomba BLID (can be obtained with discover)")
	flagSet.StringVar(&a.password, passwordArgName, "", "roomba password (can be obtained with get-password)")
	flagSet.BoolVar(&a.debug, "debug", false, "print debug logs")
	flagSet.StringVar(&a.stateLogFileName, "state-log", "", "roomba state message log file")
	flagSet.Parse(os.Args[2:])
}

func (a *App) Init() error {
	err := a.createStateLogFile()
	if err != nil {
		return err
	}

	cfg := a.createRoombaConfig()

	a.roombaClient = roomba.New(cfg)
	err = a.roombaClient.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Cleanup() error {
	if a.stateLogFile != nil {
		err := a.stateLogFile.Close()
		if err != nil {
			return err
		}
	}

	if a.roombaClient != nil {
		a.roombaClient.Disconnect()
	}

	return nil
}

func (a *App) GetClient() *roomba.Roomba {
	return a.roombaClient
}

func (a *App) createStateLogFile() error {
	if len(a.stateLogFileName) > 0 {
		f, err := os.Create(a.stateLogFileName)
		if err != nil {
			return err
		}
		a.stateLogFile = f
	}
	return nil
}

func (a *App) createRoombaConfig() *config.Config {
	cfg := config.NewFromEnv("ROOMBA_")

	cfg.Debug = a.debug

	if a.host != "" {
		cfg.Address = a.host
	}

	if a.user != "" {
		cfg.User = a.user
	}

	if a.password != "" {
		cfg.Password = a.password
	}

	cfg.StateWriter = a.stateLogFile

	return cfg
}
