// EC2 game

package main

import (
	//"errors"
	"fmt"
	"io"
	"os"

	"github.com/greenm01/ec2game/internal/core"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := parseArgs(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func parseArgs(args []string, stdout io.Writer) error {
	/*
		if len(args) < 2 {
			return errors.New("no names")
		}
		for _, name := range args[1:] {
			fmt.Fprintf(stdout, "Hi %s\n", name)
		}*/

	core.InitGame()

	return nil
}
