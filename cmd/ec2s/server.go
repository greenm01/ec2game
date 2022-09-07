package main

import (
	"errors"
	"fmt"
	"github.com/greenm01/ec2game/internal/core"
	"github.com/greenm01/ec2game/internal/server"
	"github.com/greenm01/ec2game/internal/db"
	"net"
	"time"
	"encoding/gob"
)

func initServer(path string) error {

	/* TODO: (not in order)
	 *		 1) Parse config-file
		     2) Load config from DB
			 3) Update config, write back to DB
		     4) Load latest gamestate from DB
	         5) Configure server from config
	         6) Create GameSpace and pass gamestate
			 7) Monitor system time for maintenance
	         8) Send/receive game data from user session
	*/
	
	// Load configuration file for dynamic game parameters
	c := core.Config{}
	if err := c.Load(path); err != nil {
		return err
	}

	ip := c.IP
	port := c.Port	
	maintTime := c.MaintTime
	maintPeriod := c.MaintPeriod
	
	// Load config from database
	
	dex := 0 // zero for config
	cBuff,err := db.Read(db.GenKey(dex, STARTYEAR), path)
	if err != nil { return err }
	
	var cfg core.Config
	dec := gob.NewDecoder(&cBuff)	
	err = dec.Decode(&cfg)
	if err != nil { return err }
	
	// Update game config with dynamic game settings
	cfg.IP = ip
	cfg.Port = port
	cfg.MaintTime = maintTime
	cfg.MaintPeriod = maintPeriod
	
	// Load the gamestate from database with latest year from config
	
	dex = 1 // One for gamestate
	gBuff,err := db.Read(db.GenKey(dex, cfg.GameYear), path)
	if err != nil { return err }
	
	var gs  core.GameState
	dec = gob.NewDecoder(&gBuff)	
	err = dec.Decode(&gs)
	if err != nil { return err }	
	
	// Start the game server
	
	fmt.Println("Server started.")
	
	listener, err_listen := net.Listen("tcp", ip+":"+port)
	if err_listen != nil {
		return errors.New("Game server listener failed. Exit")
	}

	fmt.Println("Server started to listen on port " + port)

	gameSpace := server.NewGameSpace(cfg, gs)
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
