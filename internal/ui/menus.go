package ui

import (
	lg "github.com/charmbracelet/lipgloss"
)

type FirstTimeMenu struct {
	Title string
	Menu string
	Help string
}

func (ftm *FirstTimeMenu) Init() {

	title := lg.NewStyle().
		Bold(true).
		SetString(" FIRST TIME MENU ").
		Foreground(lg.Color("252")).
		Background(lg.Color("21"))
	
	ftm.Title = title.String()
	
	cmdStyle := lg.NewStyle().
		Bold(true).
		Foreground(lg.Color("11"))
	
	txtStyle := lg.NewStyle().
		Foreground(lg.Color("253")).
		Width(26)
	
	h := cmdStyle.Render("H")
	h += txtStyle.Render(">elp with Commands")
	
	l := cmdStyle.Render("L")
	l += txtStyle.Render(">ist Current Empires")	

	v := cmdStyle.Render("V")
	v += txtStyle.Render(">iew Starmap")

	q := cmdStyle.Render("Q")
	q += txtStyle.Render(">uit Back to BBS")	
	
	j := cmdStyle.Render("J")
	j += txtStyle.Render(">oin this Game")

	s := cmdStyle.Render("S")
	s += txtStyle.Render(">how Game Introduction")
		
	ftm.Menu = h + l + v + q + j + s
	
	ftm.Help =	"<A> - put game into ANSI color mode\n" +
			"<H> - describe First Time Menu commands\n" +
			"<J> - join the game and control and unowned empire" +
			"<L> - list all empires in the order you specify" +
			"<Q> - quit Esterian Coquest and returns you back to BBS" +
			"<V> - view the introduction to this game"
}

type MainMenu struct {
	
}

type GeneralMenu struct {
	
}

type PlanetMenu struct {
	
}

type BuildMenu struct {
	
}

type FleetMenu struct {
	
}

type MessageMenu struct {
	
}