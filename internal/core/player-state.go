package core

import (
	"bytes"
	"encoding/gob"
)

type PlayerState struct {
	user   User
	empire Empire
	year   int
	// Other empire data with UID as key
	eNames     map[int]string
	numPlanets map[int]int
	curProd    map[int]int
	prevProd   map[int]int
}

func (p PlayerState) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(p); err != nil {
		return buff, err
	}

	return buff, nil

}