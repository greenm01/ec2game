// EC2 game

package main

import (
	//"errors"
    "fmt"
    "os"
	"io"
	
	"github.com/greenm01/ec2game/internal/core"
	
    tea "github.com/charmbracelet/bubbletea"
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
	/*
	if len(args) < 2 {
		return errors.New("no names")
	}
	for _, name := range args[1:] {
		fmt.Fprintf(stdout, "Hi %s\n", name)
	}*/

	p := tea.NewProgram(core.InitGame())
    	if err := p.Start(); err != nil {
        	fmt.Printf("Alas, there's been an error: %v", err)
        	os.Exit(1)
    	}	

	return nil
}
