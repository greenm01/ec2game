package io

import (
    bx "github.com/treilik/bubbleboxer"
	tea "github.com/charmbracelet/bubbletea"
)

type intelCmd struct {
    tui bx.Boxer
    text string
}

// satisfy the tea.Model interface
func (b intelCmd) Init() tea.Cmd                           { return nil }
func (b intelCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b intelCmd) View() string                            { return b.text }
