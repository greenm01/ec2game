package main

import "fmt"
import "github.com/speps/go-hashids/v2"

func main() {
    
    // Game database keys
    
    master := 0
    player1 := 1
    gameYear := 3000
    hd := hashids.NewData()
	hd.Salt = "Esterian Conquest"
	hd.MinLength = 6
    
	h, _ := hashids.NewWithData(hd)
	key1, _ := h.Encode([]int{master, gameYear})
    key2,_ := h.Encode([]int{player1, gameYear})
	
    fmt.Println(key1, key2)
	
    d1, _ := h.DecodeWithError(key1)
    d2, _ := h.DecodeWithError(key2)
	
    fmt.Println(d1,d2)
        
}
