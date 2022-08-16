package io

import (
    bx "github.com/treilik/bubbleboxer"
	tea "github.com/charmbracelet/bubbletea"
)

type comsCmd struct {
    tui bx.Boxer
    text string
}

func (m comsCmd) GetTui() bx.Boxer { return m.tui }
func (m comsCmd) InitBox() {}

// satisfy the tea.Model interface
func (b comsCmd) Init() tea.Cmd                           { return nil }
func (b comsCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b comsCmd) View() string                            { return b.text }
