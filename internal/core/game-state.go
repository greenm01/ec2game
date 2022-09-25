package core

import (
	"bytes"
	"encoding/gob"
)

// GameState holds the game's master data
// Starmap holds all the planet data
// Empires holds all the player data
type GameState struct {
	Full 	bool
	Year    int
	StarMap StarMap
	// Empire index is user ID
	Empires map[int]*Empire
	// User index is user's name
	Users   map[string]*User
}

func (g GameState) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(g); err != nil {
		return buff, err
	}

	return buff, nil

}
