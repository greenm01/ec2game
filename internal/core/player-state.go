package core

import (
	"bytes"
	"encoding/gob"
	"bufio"
)

type PlayerState struct {
	User   *User
	Empire *Empire
	Year   int
	GameFull bool
	// Other empire data with UID as key
	Names      map[int]string
	NumPlanets map[int]int
	CurProd    map[int]int
	PrevProd   map[int]int
}

func (p *PlayerState) Setup(gs GameState) {
	
	p.GameFull = gs.Full
	p.Year = gs.Year

	if p.User.ID != -1 {
		p.Empire = gs.Empires[p.User.ID] 
	} 

	p.Names = make(map[int]string)
	p.NumPlanets = make(map[int]int)
	p.CurProd = make(map[int]int)
	p.PrevProd = make(map[int]int)
		
	for key,e := range gs.Empires {
		p.Names[key] = e.Name
		p.NumPlanets[key] = len(e.Planets)
		p.CurProd[key] = e.CurProd
		p.PrevProd[key] = e.PrevProd
	}
	
}

func (p *PlayerState) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(p); err != nil {
		return buff, err
	}

	return buff, nil

}

func DecodePlayerState(rw *bufio.ReadWriter) (PlayerState, error) {
	
	dec := gob.NewDecoder(rw)
	var ps PlayerState
	if err := dec.Decode(&ps); err != nil {
		return ps, err
	}
	return ps,nil

}