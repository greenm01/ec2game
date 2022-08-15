package io

import (
	tea "github.com/charmbracelet/bubbletea"
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
	s := "Press q to quit!\n"
	return s
}
