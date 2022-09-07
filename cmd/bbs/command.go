package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"	
	tea "github.com/charmbracelet/bubbletea"
)

func run() error {
	
	path, err := getPath()
	if err != nil {
		return err
	}	

	p := tea.NewProgram(initialCmd())

	if err := p.Start(); err != nil {
		return err
	}
	
	fmt.Println(path)
	
	return nil
}

func initialCmd() FirstTimeCommand {
	ti := textinput.New()
	ti.Placeholder = "Esterian Conquest"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	ftm := FirstTimeCommand{textInput: ti,
		                          err: nil,}
	
	ftm.BuildMenu()
	
	return ftm
}

