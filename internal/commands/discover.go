package commands

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/szeist/roomba-go/pkg/discover"
)

func Discover() {
	discoverFlags := flag.NewFlagSet("discover", flag.ExitOnError)
	discoverTimeoutFlag := discoverFlags.Int("discover-timeout", 2, "roomba discovery timeout in seconds")
	discoverFlags.Parse(os.Args[2:])

	results, err := discover.Discover(time.Duration(*discoverTimeoutFlag) * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
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
