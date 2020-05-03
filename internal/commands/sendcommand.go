package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/szeist/roomba-go/pkg/roomba"
)

const statusWaitTimeoutMs time.Duration = 2000

func SendCommand(roombaClient *roomba.Roomba, cmd string) {
	roombaClient.WaitForStatus(statusWaitTimeoutMs)
	err := roombaClient.SendCommand(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	} else {
		fmt.Println("sent")
	}
}
