package io

import (
    bx "github.com/treilik/bubbleboxer"
	tea "github.com/charmbracelet/bubbletea"
)

type planetCmd struct {
    tui bx.Boxer
    text string
}


func (m planetCmd) GetTui() bx.Boxer { return m.tui }
func (m planetCmd) InitBox() {}

// satisfy the tea.Model interface
func (b planetCmd) Init() tea.Cmd                           { return nil }
func (b planetCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b planetCmd) View() string                            { return b.text }