package main

import (
	"errors"
	"fmt"
	"github.com/greenm01/ec2game/internal/server"
	"net"
	"time"
)

func initServer() error {

	/* TODO:
	 *		 1) Parse config-file
		     2) Load config from DB
			 3) Update config, write back to DB
		     4) Load latest gamestate from DB
	         5) Configure server from config
	         6) Create GameSpace and pass gamestate
			 7) Monitor system time for maintenance
	         8) Send/receive game data from user session
	*/
	
	fmt.Println(CServer started.")
			 7) Monitor system time for maintenance
	         8) Send/receive game data from user sessio
	port := "6666"
	
	listener, err_listen := net.Listen("tcp", ":"+port)
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
