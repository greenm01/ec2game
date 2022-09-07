package db

import (
	"bytes"
	pg "github.com/akrylysov/pogreb"
	hash "github.com/speps/go-hashids/v2"
)

const DBDIR = "db"

func Write(key string, buff bytes.Buffer, path string) error {

	path += DBDIR

	db, err := pg.Open(path, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Put([]byte(key), buff.Bytes()); err != nil {
		return err
	}

	return nil

}

func Read(key string, path string) (bytes.Buffer, error) {
	
	path += DBDIR

	var buff bytes.Buffer
	
	db, err := pg.Open(path, nil)
	if err != nil {
		return buff, err
	}
	defer db.Close()

	val, err := db.Get([]byte(key))
	if err != nil {
		return buff, err 
	}
	buff.Write(val)
	
	return buff, nil
		
}

// GenKey generates a database lookup key
// dex = 0 for game config, dex = 1 gamestate
func GenKey(dex int, gYear int) string {
	hd := hash.NewData()
	hd.Salt = "Esterian Conquest"
	hd.MinLength = 6
	h, _ := hash.NewWithData(hd)
	key, _ := h.Encode([]int{dex, gYear})
	return key
}
