package db

import (
	"bytes"
	pg "github.com/akrylysov/pogreb"
	hash "github.com/speps/go-hashids/v2"
)

func Write(key string, buff bytes.Buffer, path string) error {

	path += "db"

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
