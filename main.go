// EC2 game

package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/greenm01/ec2game/internal/ga"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no names")
	}
	for _, name := range args[1:] {
		fmt.Println("Hi",name)
	}

	// do something
	myShip := ga.Destroyer{}
	myShip.Type = 5

	return nil

}
