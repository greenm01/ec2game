package main

import (
	"errors"
	"fmt"
	"github.com/greenm01/ec2game/internal/core"
	"github.com/greenm01/ec2game/internal/db"
	"strconv"
)

type Ship = core.Ship

func newGameSetup(path string) error {

	cfg := config{}
	if err := cfg.Setup(path); err != nil {
		return err
	}

	gs := &core.GameState{Year: startYear}

	fmt.Println("\n################################")
	fmt.Println("##### Creating New EC Game #####")
	fmt.Println("################################")

	// Number of players

	np := cfg.NumPlayers()
	if np < 2 {
		e := "\nError! Minimum number of players is 2.\n" +
			"Fix the configuration file.\n"
		return errors.New(e)
	}

	/* #############################
	   ##### STARMAP & PLANETS #####
	   ############################# */

	starMap := core.StarMap{}
	if err := starMap.InitMap(np); err != nil {
		return err
	}

	/* #########################
	   ##### PLAYER SETUP  #####
	   ######################### */

	empires := make(map[int]*core.Empire)

	// Create players & assign homeworlds
	for i, hw := range starMap.HomeWorlds {
		id := i + 1
		empires[id] = &core.Empire{UID: id,
			Name:      "Rogue " + strconv.Itoa(id),
			Planets:   []int{hw},
			Tax:       50.0,
			Autopilot: true,
		}
		// Allocate resources to homeworld
		starMap.Planets[hw].InitHomeworld(id)
	}

	/*	## Assign fleets #####################################
		   	# Per classic EC, each empire get 4 starting fleets: #
		   	# - Two Fleets with one ETAC and a Cruiser escort    #
		   	# - Two Fleets with one Destroyers                   #
		   	# - Orders are to guard/blockade homeworld           #
			######################################################
	*/

	fmt.Print("...\nAssigning fleets...")
	for _, e := range empires {
		e.Fleets = make(map[int]*core.Fleet)
		// Create fleets
		for f := 0; f < 4; f++ {
			e.Fleets[f] = &core.Fleet{ID: f,
				Pos:    e.Planets[0],
				ROE:    6,
				Speed:  0,
				ETA:    0,
				Orders: 5,
			}
		}
				
		e.Fleets[0].Ships = []Ship{Ship{ID: 1, Class: 2}, Ship{ID: 2, Class: 6}}
		e.Fleets[1].Ships = []Ship{Ship{ID: 3, Class: 2}, Ship{ID: 4, Class: 6}}
		e.Fleets[2].Ships = []Ship{Ship{ID: 5, Class: 1}}
		e.Fleets[3].Ships = []Ship{Ship{ID: 6, Class: 1}}
		
	}

	fmt.Println("done!")

	gs.StarMap = starMap
	gs.Empires = empires

	fmt.Print("Initializing game database...")

	// dex = 0 for game config
	dex := 0
	buff, err := cfg.Encode()
	if err != nil {
		return err
	}

	if err := db.Write(db.GenKey(dex, startYear), buff, path); err != nil {
		return err
	}

	// dex = 1 for gamestate
	dex = 1
	buff, err = gs.Encode()
	if err != nil {
		return err
	}
	if err := db.Write(db.GenKey(dex, startYear), buff, path); err != nil {
		return err
	}

	fmt.Println("done!")

	fmt.Println("Launch date & time set to", cfg.LaunchDate(), "@", cfg.MaintTime())

	return nil

}
