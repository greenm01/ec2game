package io

import (
    bx "github.com/treilik/bubbleboxer"
	tea "github.com/charmbracelet/bubbletea"
)

type reportCmd struct {
    tui bx.Boxer
    text string
}

// satisfy the tea.Model interface
func (b reportCmd) Init() tea.Cmd                           { return nil }
func (b reportCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b reportCmd) View() string                            { return b.text }
