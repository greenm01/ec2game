package bbsui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/bubbles/spinner"
	lg "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var expert = false

// First Time Menu
type FirstTime struct {
	err error
	showhelp bool
	title string
	menu string
	help string
	cmdLine string
	spin  spinner.Model
}

func (ft *FirstTime) Build() {

	title := lg.NewStyle().
		Bold(true).
		SetString(" FIRST TIME MENU ").
		Foreground(lg.Color("252")).
		Background(lg.Color("21"))
	
	ft.title = title.String()
	
	// Yellow
	yellow := lg.NewStyle().
		Bold(true).
		Foreground(lg.Color("11"))
	
	// White
	white := lg.NewStyle().
		Foreground(lg.Color("253")).
		Width(26)
	
	h := yellow.Render("H")
	h += white.Render(">elp with Commands")
	
	l := yellow.Render("L")
	l += white.Render(">ist Current Empires")	

	v := yellow.Render("V")
	v += white.Render(">iew Starmap")

	q := yellow.Render("Q")
	q += white.Render(">uit Back to BBS")	
	
	j := yellow.Render("J")
	j += white.Render(">oin this Game")

	s := yellow.Render("S")
	s += white.Render(">how Game Introduction")
		
	mtxt := h + l + v + q + j + s
	
	b := newBoxWithLabel()
	ft.menu = b.Render(ft.title, mtxt, 80)
	
	cmd := lg.NewStyle().
		Bold(true).
		SetString("FIRST TIME COMMAND").
		Foreground(lg.Color("252")).
		Background(lg.Color("21"))
	
	txt := yellow.Render("H Q L J V S")
	lArrow := strings.TrimSpace(white.Render(" <-"))
	rArrow := strings.TrimSpace(white.Render("->"))
	
	ft.cmdLine = cmd.String() + lArrow + txt + rArrow + " "

	// Yellow on orange
	yoo := lg.NewStyle().
		Bold(true).
		Foreground(lg.Color("11")).
		Background(lg.Color("202"))
	
	ft.help = yoo.Render(" <V> - View the game's Starmap \n" +
	" <H> - Describe First Time Menu commands \n" + 
	" <J> - Join the game and control an unowned empire \n" +
	" <L> - List all empires in the order you specify \n" +
	" <Q> - Quit Esterian Coquest and return back to BBS \n" +
	" <S> - Show the introduction to this game ")
	
	ft.spin = spinner.New()
	ft.spin.Spinner = spinner.Dot
	ft.spin.Style = lg.NewStyle().Foreground(lg.Color("205"))

}

func (ft FirstTime) Init() tea.Cmd {
	// clear the screen with ANSI code
	fmt.Print("\033[H\033[2J")	
	return ft.spin.Tick
}

func (ft FirstTime) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlQ, tea.KeyCtrlC, tea.KeyEsc:
			return ft, tea.Quit
		case tea.KeyRunes:
			switch string(msg.Runes) {
				case "q":
					return ft, tea.Quit
				case "h":
					ft.showhelp = !ft.showhelp
					return ft, cmd
			}
		}
	case error:
		ft.err = msg
		return ft, nil
	default:
		var cmd tea.Cmd
		ft.spin, cmd = ft.spin.Update(msg)
		return ft, cmd	
	}

	return ft, cmd

}

func (ft FirstTime) View() string {

	var s strings.Builder
	s.WriteString(ft.menu + "\n\n" + ft.cmdLine + ft.spin.View())
	if(ft.showhelp) {
		s.WriteString("\n\n" + ft.help)
	}
	return s.String()

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