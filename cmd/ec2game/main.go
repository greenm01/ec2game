// EC2 game

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"log"

	"github.com/rivo/tview"

	"github.com/greenm01/ec2game/internal/core"
	"github.com/greenm01/ec2game/internal/ui"
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
		fmt.Fprintf(stdout, "Hi %s", name)
	}

	// Create and initialize the game
	game := &core.EC2{
		TView:      tview.NewApplication(),
		PageHolder: tview.NewPages(),
	}

	game.Initialise()
	
	ui.SetUniversalHandlers(game)
	
	// Run the app.
	log.Println("Running app...")
	if err := game.TView.Run(); err != nil {
		log.Println(err)
	}	

	// Shut it all down
	game.Shutdown()

	return nil
}
