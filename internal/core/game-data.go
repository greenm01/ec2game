package core

import (
	"encoding/gob"
	"os"
	"sync"
	"time"
)

// Make code safe for concurrent use
var lock sync.Mutex

type GameData struct {
	players      []string
	gameInitTime time.Time
	gameYear     int
	dbKeys       []string
}

func saveGameData(filePath string, gd GameData) error {

	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(filePath)
	defer file.Close()
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(gd)
	}
	return err
}

func loadGameData(filePath string, gd GameData) error {

	lock.Lock()
	defer lock.Unlock()

	file, err := os.Open(filePath)
	defer file.Close()
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(gd)
	}
	return err
}
