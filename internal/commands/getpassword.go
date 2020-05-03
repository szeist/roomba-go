package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/szeist/roomba-go/pkg/roombapass"
)

func GetPassword() {
	getPasswordFlags := flag.NewFlagSet("get-password", flag.ExitOnError)
	getPasswordHostFlag := getPasswordFlags.String("host", "", "roomba ip address")
	getPasswordFlags.Parse(os.Args[2:])

	if len(*getPasswordHostFlag) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify roomba address to get password from!\n")
		getPasswordFlags.Usage()
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Press the home button on the roomba for at least 2 seconds.\nPress enter when ready!")
	fmt.Scanln()
	password, err := roombapass.GetPassword(*getPasswordHostFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "The roomba password is:")
		fmt.Println(password)
	}
}
