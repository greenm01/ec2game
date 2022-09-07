package core

import (
	"bytes"
	"encoding/gob"
)

type PlayerState struct {
	User   User
	Empire Empire
	Year   int
	// Other empire data with UID as key
	ENames     map[int]string
	NumPlanets map[int]int
	CurProd    map[int]int
	PrevProd   map[int]int
}

func (p PlayerState) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(p); err != nil {
		return buff, err
	}

	return buff, nil

}