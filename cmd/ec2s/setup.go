package main

import (
	"github.com/greenm01/ec2game/internal/core"
)

/* TODO: 1) Generate Starmap based on # of players
         2) Setup NPC rogue empires
	     3) Save data to databse (create db package)
		 4) Place server game-state in standby, waiting for 
            new players to join
*/

func NewGameSetup(config ConfigData) error {
	sm := core.StarMap{}
	return sm.Init(4)
}
