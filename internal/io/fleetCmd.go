package io

import (
	tea "github.com/charmbracelet/bubbletea"
    bx "github.com/treilik/bubbleboxer"
)

type fleetCmd struct {
    tui bx.Boxer
    text string
}

// satisfy the tea.Model interface
func (b fleetCmd) Init() tea.Cmd                           { return nil }
func (b fleetCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b fleetCmd) View() string                            { return b.text }
