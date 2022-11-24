package core

import (
	"bytes"
	"encoding/gob"
)

// Empire contains empire data
type Empire struct {
	ID        int
	Name      string
	Planets   []int
	PrevPlanets int
	Fleets    map[int]*Fleet
	Reports   []Report
	Messages  []Message
	Tax       int
	PDB       PlanetDB
	Autopilot bool
	Status    string
	CurProd   int
	PrevProd  int
	Bio 	  string
}

// NewEmpire contains info to creae a new empire from client
type NewEmpire struct {
	Name string
	Bio string
}

func (e NewEmpire) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(e); err != nil {
		return buff, err
	}

	return buff, nil

}