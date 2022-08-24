package main

import(
    "fmt"
    "net"
    "os"
    "time"
	"log"
	"strings"
	
    "github.com/greenm01/ec2game/internal/server"
	
	"github.com/kylelemons/go-gypsy/yaml"   
)

const configFile = "config.yaml"

func cmdLineError() {
	usage := "\nEsterian Conquest v2.0 GAME SERVER\n\n" +
	         "Usage: ec2s <command> [game path]\n\n" +
	         "The commands are:\n\n" +
	         "                      new            Initialize a new game\n" +
	   	     "                      run            Start the game server\n" +
			 "                      maint          Manually run turn maintenance\n" +
			 "                      stats          Display game stastics\n\n" +
	         "- Be sure to specify the game folder directory, e.g. ec2s new /User/mag/ec2/game1\n\n" +
			 "- Ensure you drop an updated config.yaml for each new game in this folder\n" +
		     "  > Example config.yaml, with required fields:\n" +
			 "  > \n" + 
		     "  > players: 4                       # Number of platers in game\n" +
	         "  > host: Toys In The Attic BBS      # Host system name\n" +
			 "  > sysop: Mason Austin Green        # System operator name\n" +
			 "  > launchDate: 2022-08-23           # Day to officialy start the game: YEAR-MM-DD\n" + 
			 "  > maintPeriod: 24                  # Time between maintenance runs (hours) \n" +
			 "  > ipaddress: localhost             # Your server's IP address\n" +
			 "  > port: 7777                       # Port number\n\n" +
		     "- To delete a game, delete the folder (save the config.yaml file for later use)\n\n"
	
	fmt.Println(usage)
}

func main() {
	
	// Verify that a subcommand has been provided
    // os.Arg[0] is the main command
	// os.Arg[1] is the sub command
    if len(os.Args) != 3 {
        cmdLineError()
        os.Exit(1)
    }
	
	path := strings.TrimSpace(os.Args[2])
	
	switch os.Args[1] {
    	case "new":
			newGame(path)		
	    case "run":
	        runGame(path)
	    default:
			cmdLineError()
	        os.Exit(1)
    }	

	os.Exit(0)	
}

func newGame(path string){
	
	filePath := path+configFile	
	config, err := yaml.ReadFile(filePath)
	if err != nil {
		log.Fatalf("readfile(%q): %s", filePath, err)
	}
	
	host, _ := config.Get("host")
	fmt.Println(host)

}

func runGame(path string){
	initServer()
}

func initServer() {
  
    fmt.Println("Server started.")

	port := "6666"
	listener, err_listen := net.Listen("tcp", ":" + port)
	if err_listen != nil {
		fmt.Println("Game server listening failed. Exit.")
		os.Exit(1)
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
    
}