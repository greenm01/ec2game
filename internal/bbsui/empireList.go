package bbsui

import (
	"strconv"
	"strings"
	"github.com/greenm01/ec2game/internal/core"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/table"
	lg "github.com/charmbracelet/lipgloss"	
)

var baseStyle = lg.NewStyle().
	BorderStyle(lg.NormalBorder()).
	BorderForeground(lg.Color("240"))

type EmpireList struct {
	table table.Model
	year string
	footer string
	bios map[string]string
}

func (e EmpireList) Init() tea.Cmd {
	return nil
}

func (e *EmpireList) Build(ps core.PlayerState) {
	

	title := lg.NewStyle().
		Bold(true).
		SetString("STARDATE ").
		Foreground(lg.Color("205"))

	yellow := lg.NewStyle().
		Bold(true).
		Foreground(lg.Color("11"))
	
	e.year = " " + title.String() + yellow.Render(strconv.Itoa(ps.Year)) + "\n"
	
	e.footer = lg.NewStyle().
		SetString(" (esc to quit) - previous year shown in parens.").
		Foreground(lg.Color("57")).String()

	columns := []table.Column{
		{Title: "Empire Name", Width: 25},
		{Title: "ID", Width: 6},
		{Title: "Planets Owned", Width: 15},
		{Title: "Current Production", Width: 20},
		{Title: "Status", Width:7},
	}
	
	e.bios = make(map[string]string)
	
	var rows []table.Row	
	for id,empire := range ps.Names {
		planets := strconv.Itoa(ps.NumPlanets[id]) +
		           " (" + strconv.Itoa(ps.PrevPlanets[id]) + ")"
		prod := strconv.Itoa(ps.CurProd[id]) + 
		        " (" + strconv.Itoa(ps.PrevProd[id]) + ")"
		i := strconv.Itoa(id)
		r := table.Row{empire, i, planets, prod, ps.Status[id]}
		rows = append(rows,r)
		e.bios[i] = ps.Bios[id]
	}
	
	l := len(ps.Names)
	if l > 24 { l = 24}
	
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(l),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lg.NormalBorder()).
		BorderForeground(lg.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lg.Color("229")).
		Background(lg.Color("57")).
		Bold(false)
	t.SetStyles(s)
	
	e.table = t
	
}

func (e EmpireList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return arb.GetModel("ftm"), changeMenu("ftm")
		/*case "enter":
			return tea.Batch(
				tea.Printf("Let's go to %s!", e.table.SelectedRow()[1]),
			)*/
		}
	}
	var cmd tea.Cmd
	e.table, cmd = e.table.Update(msg)
	return  e, cmd

}

func (e EmpireList) View() string {
	var s strings.Builder
    bio := e.bios[e.table.SelectedRow()[1]]
	render := baseStyle.Render(e.table.View()) + "\n"	
	s.WriteString(e.year + render + e.footer + "\n\n " + bio)
	return s.String()
}






