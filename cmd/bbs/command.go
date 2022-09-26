package main

import (
	"github.com/greenm01/ec2game/internal/core"
	ui "github.com/greenm01/ec2game/internal/bbsui"
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

// ftm initializes the First Time Menu
func ftm() ui.Menu {
	var menu ui.Menu
	f := ui.FirstTime{}
	f.Build()
	menu.Build("ftm", &f)
	pager := ui.Pager{}
	pager.Build(ui.Intro())
	menu.Build("intro", &pager)
	return menu
}

