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

	p := tea.NewProgram(initialCmd())

	if err := p.Start(); err != nil {
		return err
	}
	
	return nil
}

func initialCmd() ui.FirstTime {
	ft := ui.FirstTime{}
	ft.Build()
	return ft
}

