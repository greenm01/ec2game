package main

import (
	"github.com/greenm01/ec2game/internal/core"
)

func run() error {
	
	/* TODO: load playerstate from server */
	
	dropPath, gamePath, err := getPaths("drop", "game")
	if err != nil { 
		return err
	}
	
	config := core.Config{}
	config.Load(gamePath)
	
	alias,_,_,_ := dropFileData(dropPath)
	
	bbs := bbsClient{user:alias,ip:config.IP,port:config.Port}
	return bbs.Run() 
}



