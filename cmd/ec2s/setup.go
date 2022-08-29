package main

import (
	"errors"
	"fmt"
	"github.com/greenm01/ec2game/internal/core"
	"strconv"
	"time"
)

type Ship = core.Ship

/* TODO: * Save data to databse (create db package)
		 * Place server game-state in standby, waiting for
            new players to join
*/

func newGameSetup(config map[string]interface{}) error {

	fmt.Println("\n############################")
	fmt.Println("#### Creating New Game #####")
	fmt.Println("############################")

	// Number of players
	
	np := config["players"].(int)
	if np < 2 {
		e := "\nError! Minimum number of players is 2.\n" +
			"Fix the configuration file.\n"
		return errors.New(e)
	}

	/* #############################
	   ##### STARMAP & PLANETS #####
	   ############################# */

	starMap := &core.StarMap{}
	if err := starMap.InitMap(np); err != nil {
		return err
	}

	/* #########################
	   ##### PLAYER SETUP  #####
	   ######################### */

	empires := make(map[int]*core.Empire)

	// Create players & assign homeworlds 
	for i, hw := range starMap.HomeWorlds {
		empires[i] = &core.Empire{UID: i,
			Name: "Rogue " + strconv.Itoa(i),
			Planets:    []int{hw},
			TaxRate:    50.0,
		}
		// Allocate resources to homeworld
		starMap.Planets[hw].InitHomeworld(i)
	}

	/*	## Assign fleets #####################################
	   	# Per classic EC, each empire get 4 starting fleets: #
	   	# - Two Fleets with one ETAC and a Cruiser escort    #
	   	# - Two Fleets with one Destroyers                   #
	   	# - Orders are to guard/blockade homeworld           #
		###################################################### 
	*/ 

	fmt.Print("...\nAssigning fleets...")
	for _, p := range empires {
		p.Fleets = make(map[int]*core.Fleet)
		// Create fleets
		for f := 0; f < 4; f++ {
			p.Fleets[f] = &core.Fleet{ID: f,
				Pos:    p.Planets[0],
				ROE:    6,
				Speed:  0,
				ETA:    0,
				Orders: 5,
			}
		}

		// Assign starting ships
		p.Fleets[0].Ships = []Ship{Ship{ID: 1, Class: 2},
			Ship{ID: 2, Class: 6}}

		p.Fleets[1].Ships = []Ship{Ship{ID: 3, Class: 2},
			Ship{ID: 4, Class: 6}}

		p.Fleets[2].Ships = []Ship{Ship{ID: 5, Class: 1}}

		p.Fleets[3].Ships = []Ship{Ship{ID: 6, Class: 1, AR:2}}
	}
	
	fmt.Println("done!")

	gs := &gameState{LaunchTime: config["launchTime"].(time.Time), 
		             StarMap: starMap, Empires: empires}

	fmt.Println("Launch date & time set to ",gs.LaunchTime)
	
	return nil
}
