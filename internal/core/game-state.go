package core

import (
	"bytes"
	"encoding/gob"
)

// GameState holds the game's master data
// Starmap holds all the planet data
// Empires holds all the player data
type GameState struct {
	Year    int
	StarMap StarMap
	Empires map[int]*Empire
	Users   []User
}

func (g GameState) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(g); err != nil {
		return buff, err
	}

	return buff, nil

}
