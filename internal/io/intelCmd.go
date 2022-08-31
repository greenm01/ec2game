package io

import (
	tea "github.com/charmbracelet/bubbletea"
	bx "github.com/treilik/bubbleboxer"
)

type intelCmd struct {
	tui  bx.Boxer
	text string
}

func (m intelCmd) GetTui() bx.Boxer { return m.tui }
func (m intelCmd) InitBox()         {}

// satisfy the tea.Model interface
func (b intelCmd) Init() tea.Cmd                           { return nil }
func (b intelCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b intelCmd) View() string                            { return b.text }
