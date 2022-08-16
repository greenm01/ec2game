package io

import (
	"fmt"
	
    bx "github.com/treilik/bubbleboxer"
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

type GameFrame struct {
	tabs map[string]*bx.Boxer
	tui *bx.Boxer
}

func (m *GameFrame) InitCmd() {
	m.tabs = map[string]*bx.Boxer{
		planetBox: &bx.Boxer{},
		fleetBox:  &bx.Boxer{},
		intelBox:  &bx.Boxer{},
		reportBox: &bx.Boxer{},
		comsBox:   &bx.Boxer{},
	}

	initReportCmd(m.tabs[reportBox])
	fmt.Println(m.tabs[reportBox].LayoutTree.GetAddress())
	m.tui = m.tabs[reportBox]
	
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
	
	case tea.WindowSizeMsg:
		m.tui.UpdateSize(msg)	
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
    	BorderStyle(lpg.RoundedBorder()).
    	BorderForeground(lpg.Color("34")).
    	BorderBackground(lpg.Color("0")).
    	BorderTop(true).
    	BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		SetString("Esterian Conquest")

	var s = style.Render(m.tui.View())
	s += "\nCtrl-C to Quit"
	return s
}
