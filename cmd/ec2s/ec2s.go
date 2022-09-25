package main

/* TODO: 1) Enable new players to join before game launch */

import (
	"errors"
	"fmt"
	"os"
	"strings"
	
	"github.com/greenm01/ec2game/internal/global"
	
)

const STARTYEAR = global.STARTYEAR 
const DBDIR = global.DBDIR

const (
	exitFail   = 1
	usage      = "\nEsterian Conquest v2.0 GAME SERVER\n\n" +
		"Usage: ec2s <command> [game path]\n\n" +
		"The commands are:\n\n" +
		"                      new          # Initialize a new game\n" +
		"                      run          # Start the game server\n" +
		"                      maint        # Manually run turn maintenance [TODO]\n" +
		"                      stats        # Display game stastics [TODO]\n\n" +
		"- Be sure to specify the game folder directory, e.g. ec2s new /User/mag/ec2/game1/\n\n" +
		"- Ensure you drop an updated config.nt (nestedtext.org format) for each new game in this folder\n" +
		"  > Example config.nt, with required fields:\n" +
		"  > \n" +
		"  > players: 4                     # Number of players in game\n" +
		"  > host: Toys In The Attic BBS    # Host system name\n" +
		"  > sysop: Mason Austin Green      # System operator name\n" +
		"  > launchDate: 2022-08-23         # Day to officialy start the game: YEAR-MM-DD\n" +
		"  > maintPeriod: 24                # Time between maintenance runs (hours) \n" +
		"  > maintTime: 00:01               # Daily maintenance time (hh:mm) 24hr format\n" +
		"  > ip: localhost                  # Your server's IP address\n" +
		"  > port: 1992                     # Port number\n\n" +
		"- To delete a game, delete the folder (save the config.nt file for later use)\n"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string) error {
	if len(args) != 3 {
		return errors.New(usage)
	}

	path := strings.TrimSpace(args[2])

	switch args[1] {
	case "new":
		return newGame(path)
	case "run":
		return runGame(path)
	default:
		return errors.New(usage)
	}

	return nil
}

func newGame(path string) error {

	if _, err := os.Stat(path + DBDIR); !os.IsNotExist(err) {
		// path/to/whatever/ exists
		return errors.New("Error: game database already exists in this location.")
	}

	if err := newGameSetup(path); err != nil {
		return err
	}

	return nil

}

func runGame(path string) error {

	// Load game config and gamestate
	// Init server
	// If year = 3,000 AND user not in game, then show first-time-menu
	// If year = 3,000 and user in game, show regular menu

	if _, err := os.Stat(path + DBDIR); os.IsNotExist(err) {
		// path/to/whatever/ exists
		return errors.New("Error: game does not exist in specified path.")
	}
	
	return initServer(path)
}
