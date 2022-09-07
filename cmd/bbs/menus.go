package main

import (
	"fmt"
	"strings"
	
	"github.com/greenm01/ec2game/internal/ui"
	
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	
)

type FirstTimeCommand struct {
	
	textInput textinput.Model
	err       error

	menu string
}

func (ftc *FirstTimeCommand) BuildMenu() {

	b := ui.NewBoxWithLabel()
	
	ftm := ui.FirstTimeMenu{}
	ftm.Init()
	
	ftc.menu = b.Render(ftm.Title,ftm.Menu,80)

}

func (ftc FirstTimeCommand) Init() tea.Cmd {
	
	// clear the screen with ANSI code
	fmt.Print("\033[H\033[2J")	
	return textinput.Blink
	
}

func (ftc FirstTimeCommand) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return ftc, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		ftc.err = msg
		return ftc, nil
	}

	ftc.textInput, cmd = ftc.textInput.Update(msg)
	return ftc, cmd

}

func (ftc FirstTimeCommand) View() string {


	
	var s strings.Builder
	s.WriteString(ftc.menu)
	s.WriteString(fmt.Sprintf("\nStanding by for orders, Admiral:\n\n%s\n\n%s",
		ftc.textInput.View(), "(esc to quit)\n"))
	return s.String()

}


