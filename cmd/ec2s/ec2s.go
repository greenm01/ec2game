package main

/* TODO: 1) Consider switching config file to 
            NestedText format
		 2) Setup game data, save to database
         3) Enable new players to join before game launch

  NOTES:

		 1) For the BBS client, consider sticking with the 
            classic EC game menu formating, ANSI compliant.
	        This will anable playtesting of the core game 
		    without spending too much time developing a nice 
            "modern" UI
         2) Once the core game is stable, donsider a new command 
            dashboard client for a terminal based game (not BBS)
            that is UTF-8 compliant.
*/

import(
    "fmt"
    "net"
    "os"
    "time"
	"strings"
	"errors"
	    
	"github.com/greenm01/ec2game/internal/server"
	
)

const (
	
	configFile = "config.nt"
 	exitFail = 1
	usage = "\nEsterian Conquest v2.0 GAME SERVER\n\n" +
	         "Usage: ec2s <command> [game path]\n\n" +
	         "The commands are:\n\n" +
	         "                      new            Initialize a new game\n" +
	   	     "                      run            Start the game server\n" +
			 "                      maint          Manually run turn maintenance [TODO]\n" +
			 "                      stats          Display game stastics [TODO]\n\n" +
	         "- Be sure to specify the game folder directory, e.g. ec2s new /User/mag/ec2/game1\n\n" +
			 "- Ensure you drop an updated config.yaml for each new game in this folder\n" +
		     "  > Example config.yaml, with required fields:\n" +
			 "  > \n" + 
		     "  > players: 4                       # Number of players in game\n" +
	         "  > host: Toys In The Attic BBS      # Host system name\n" +
			 "  > sysop: Mason Austin Green        # System operator name\n" +
			 "  > startDate: 2022-08-23            # Day to officialy start the game: YEAR-MM-DD\n" + 
			 "  > maintPeriod: 24                  # Time between maintenance runs (hours) \n" +
			 "  > ip: localhost                    # Your server's IP address\n" +
			 "  > port: 7777                       # Port number\n\n" +
		     "- To delete a game, delete the folder (save the config.yaml file for later use)\n"
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
			
	config, err := loadConfig(path); 
	if err != nil { return err	}

	return newGameSetup(config)
}

func runGame(path string) error {
	return initServer()
}

func initServer() error {
  
    fmt.Println("Server started.")

	port := "6666"
	listener, err_listen := net.Listen("tcp", ":" + port)
	if err_listen != nil {
		return errors.New("Game server listener failed. Exit")
	}
	
	fmt.Println("Server started to listen on port " + port)

	gameSpace := server.NewGameSpace()
	// listen
	gameSpace.Listen()

	for {
		conn, err_ac := listener.Accept()
		if err_ac != nil {
			fmt.Println("Connection accepting failed.")
			conn.Close()
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("A new connection accepted.")
		gameSpace.Connect(conn)  
	}  
 
	return nil   
}

