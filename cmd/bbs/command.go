package main

import (
	ui "github.com/greenm01/ec2game/internal/bbsui"
	tea "github.com/charmbracelet/bubbletea"
)

func run() error {
	/* TODO: load playerstate from server
	path, err := getPath()
	if err != nil {
		return err
	}*/	
	ui.Cls()
	p := tea.NewProgram(initialCmd())
	if err := p.Start(); err != nil {
		return err
	}
	ui.Cls()
	return nil
}

func initialCmd() ui.Menu {
	var menu ui.Menu
	ftm := ui.FirstTime{}
	ftm.Build()
	menu.Build("ftm", &ftm)
	pager := ui.Pager{}
	pager.Build(ui.Intro())
	menu.Build("intro", &pager)
	return menu
}

