package main

import (
	"fmt"
	"errors"
	"strconv"
	"github.com/greenm01/ec2game/internal/core"
)

/* TODO: 1) Generate Starmap based on # of players
         2) Setup NPC rogue empires
	     3) Save data to databse (create db package)
		 4) Place server game-state in standby, waiting for 
            new players to join
*/

func NewGameSetup(config ConfigData) error {

	// Number of players
    np,err := strconv.Atoi(config.Players)
    
    if err != nil || np < 2 {
        e := "\nError! Minimum number of players is 2.\n" +
             "Fix the configuration file.\n"
        return errors.New(e)
    }	

    /* #############################
       ##### STARMAP & PLANETS #####
       ############################# */   	
	
	fmt.Println("\n############################")
	fmt.Println("#### Creating New Game #####")
	fmt.Println("############################")
	
	starMap := core.StarMap{}
	starMap.InitMap(np)

	/* #########################
       ##### PLAYER SETUP  #####
       ######################### */   	
	
	players := make(map[int]*core.Player)

	for i,hw := range starMap.HomeWorlds {
		players[i] = &core.Player{UID:i}
		name := "Rogue " + strconv.Itoa(i)
		players[i].InitPlayer(name, hw)
	}		
	
	return err
}
