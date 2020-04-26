package roomba

import (
	"errors"
	"time"
)

var commands = []string{"start", "clean", "pause", "stop", "resume", "dock", "evac", "train"}

type command struct {
	Command   string `json:"command"`
	Time      int64  `json:"timestamp"`
	Initiator string `json:"initiator"`
}

func createCommand(cmd string) (*command, error) {
	if !isValidCommand(cmd) {
		return nil, errors.New("Invalid command: " + cmd)
	}

	return &command{
		Command:   cmd,
		Time:      time.Now().UnixNano() / 1000000,
		Initiator: "localApp",
	}, nil
}

func isValidCommand(cmd string) bool {
	for _, item := range commands {
		if cmd == item {
			return true
		}
	}
	return false
}
