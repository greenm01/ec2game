package io

import (
	//"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lpg "github.com/charmbracelet/lipgloss"
)

const (
	planetBox = "planets"
	fleetBox = "fleets"
	intelBox = "intel"
	reportBox = "reports"
	comsBox = "coms"
)

type cmdTab interface{}
	
type GameFrame struct {
	tabs map[string]cmdTab
}

func (m *GameFrame) InitCmd() {
	m.tabs = map[string]cmdTab{
		planetBox: planetCmd{},
		fleetBox:  fleetCmd{},
		intelBox:  intelCmd{},
		reportBox: reportCmd{},
		comsBox:   comsCmd{},
	}

}

func (m GameFrame) Init() tea.Cmd { 
	
	return nil
}

func (m GameFrame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m GameFrame) View() string {
	//s := "Press q to quit!\n"
	/*
	var style  = lpg.NewStyle().
    	Bold(true).
    	Foreground(lpg.Color("#FAFAFA")).
    	Background(lpg.Color("#7D56F4")).
    	PaddingTop(2).
    	PaddingLeft(4).
    	Width(22)	
	*/
	// Set a rounded, yellow-on-purple border to the top and left
	var style = lpg.NewStyle().
		Width(90).
		Height(20).
    	BorderStyle(lpg.RoundedBorder()).
    	BorderForeground(lpg.Color("34")).
    	BorderBackground(lpg.Color("0")).
    	BorderTop(true).
    	BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		SetString("Esterian Conquest")

	var s = style.Render("Hello, Commander!")
	s += "\nCtrl-C to Quit"
	return s
}
