package core

import (
    "time"
    "sync"
    "encoding/gob"
)

// Make code safe for concurrent use
var lock sync.Mutex

type GameData struct {
    players []string
    gameInitTime time.Time
    gameYear int
    dbKeys []string       
}

func saveGameData(filePath string, gd GameData) error {
    
    lock.Lock()
    defer lock.Unlock()    
    
    file, err := os.Create(filePath)
    defer file.close()
    if err == nil {
        encoder := gob.NewEncoder(file)
        err = encoder.Encode(gd)
    }
    return err
}

func loadGameData(filePath string, gd GameData) error {
    
    lock.Lock()
    defer lock.Unlock()
    
    file, err := os.Open(filePath)i
    defer file.Close()
    if err == nil {
        decoder := gob.NewDecoder(file)
        err = decoder.Decode(gd)
    }
    return err   
}
