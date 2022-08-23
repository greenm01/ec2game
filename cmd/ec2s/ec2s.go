package main

import(
    "fmt"
    "net"
    "os"
    "time"
	"flag"
	"log"
	"strings"
	
    "github.com/greenm01/ec2game/internal/server"
	
	"github.com/kylelemons/go-gypsy/yaml"   
)

const configFile = "config.yaml"

func main() {
	
	//https://www.rapid7.com/blog/post/2016/08/04/build-a-simple-cli-tool-with-golang/	// Subcommands
	// TODO: https://lightstep.com/blog/getting-real-with-command-line-arguments-and-goflags    
	
	newGameCmd := flag.NewFlagSet("new", flag.ExitOnError)
    runGameCmd := flag.NewFlagSet("run", flag.ExitOnError)	

	pathNew := newGameCmd.String("path", "", "Game directory path")
	pathRun := runGameCmd.String("path", "", "Game directory path")
	var path string
		
	// Verify that a subcommand has been provided
    // os.Arg[0] is the main command
	// os.Arg[1] is the sub command
    if len(os.Args) < 2 {
        fmt.Println("'new' or 'run' command is required")
        os.Exit(1)
    }	
	
	switch os.Args[1] {
    	case "new":
        	newGameCmd.Parse(os.Args[2:])
	    case "run":
	        runGameCmd.Parse(os.Args[2:])
	    default:
	        flag.PrintDefaults()
	        os.Exit(1)
    }	
	
	if newGameCmd.Parsed() {
		if len(*pathNew) == 0 {
			newGameCmd.PrintDefaults()
			os.Exit(1)
		}
		path = *pathNew	
	} else if runGameCmd.Parsed() {
		if len(*pathRun) == 0 {
			runGameCmd.PrintDefaults()
			os.Exit(1)
		}
		path = *pathRun		
	} else {
		// error
	}
	
	
	//command := os.Args[0]
	
	filePath := strings.TrimSpace(path)+configFile	
	config, err := yaml.ReadFile(filePath)
	if err != nil {
		log.Fatalf("readfile(%q): %s", filePath, err)
	}
	
	welcome, _ := config.Get("welcome")
	fmt.Println(welcome)

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