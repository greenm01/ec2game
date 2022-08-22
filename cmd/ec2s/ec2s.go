package main

import(
    "fmt"
    "net"
    "os"
    "time"
	
    "github.com/greenm01/ec2game/internal/server"
    
)

func main() {
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