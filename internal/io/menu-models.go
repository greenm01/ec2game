package io

import (
	"fmt"
	"strings"
    //"golang.org/x/text/encoding/charmap"	
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
	tabs map[string]bx.Boxer
	tui bx.Boxer
	
	frame lpg.Style
	cmdLine string
}

func (m *GameFrame) InitCmd() {
	m.tabs = map[string]bx.Boxer{
		planetBox: bx.Boxer{},
		fleetBox:  bx.Boxer{},
		intelBox:  bx.Boxer{},
		reportBox: bx.Boxer{},
		comsBox:   bx.Boxer{},
	}

	m.tui = m.tabs[reportBox]
	initReportCmd(&m.tui)
	m.drawFrame()
	
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
		fmt.Println("Terminal width = ", msg.Width)	
		//m.tui.UpdateSize(msg)	
		fmt.Println("Boxer width = ", m.tui.LayoutTree.GetWidth())
		
	}

	return m, nil
}

func (gf *GameFrame) drawFrame() {
	
	// white on black
	var wBlack = lpg.NewStyle().
		Bold(true).
		Background(lpg.Color("0")).
		Foreground(lpg.Color("15"))

	// yellow on blue
	var yBlue = lpg.NewStyle().
		Bold(true).
		Background(lpg.Color("#0000ff")).
		Foreground(lpg.Color("#ffff00"))
	
	// white on blue
	var wBlue = lpg.NewStyle().
		Background(lpg.Color("#0000ff")).
		Foreground(lpg.Color("#ffffff"))

	sep := wBlack.Render(" ")
	
	gf.cmdLine = wBlack.Render("\nCtrl + ") +
		wBlue.Render(" <") + yBlue.Render("r") +
		wBlue.Render(">") + yBlue.Render(" REPORTS ") + sep +
		wBlue.Render(" <") + yBlue.Render("p") +
		wBlue.Render(">") + yBlue.Render(" PLANETS ") + sep +
		wBlue.Render(" <") + yBlue.Render("f") +
		wBlue.Render(">") + yBlue.Render(" FLEETS ") + sep +
		wBlue.Render(" <") + yBlue.Render("i") +
		wBlue.Render(">") + yBlue.Render(" INTEL ") + sep +
		wBlue.Render(" <") + yBlue.Render("c") +
		wBlue.Render(">") + yBlue.Render(" COMS ") + sep +
		wBlue.Render(" <") + yBlue.Render("q") +
		wBlue.Render(">") + yBlue.Render(" QUIT ") 
	
	// Set a rounded, yellow-on-purple border to the top and left
	gf.frame = lpg.NewStyle().
		Width(132).
		Height(37).
    	BorderStyle(lpg.RoundedBorder()).
    	BorderForeground(lpg.Color("34")).
    	BorderBackground(lpg.Color("0")).
    	BorderTop(true).
    	BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
	
}

func (m GameFrame) View() string {
	
	var s strings.Builder
	s.WriteString(m.frame.Render(m.tui.View()))
	s.WriteString(m.cmdLine)
	return s.String()
	
}
