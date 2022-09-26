package bbsui

import (
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
}

func (e *EmpireList) Build(ps core.PlayerState) {
	
	columns := []table.Column{
		{Title: "Empire Name", Width: 25},
		{Title: "ID", Width: 6},
		{Title: "Planets Owned", Width: 15},
		{Title: "Current Production", Width: 20},
		{Title: "Status", Width:7},
	}
	
	
	rows := []table.Row{
		{"The Zarkonian Empire", "1", "1", "100","ALIVE"},
		{"Zzzzrrr", "2", "1", "100","ALIVE"},
		{"Master Blaster", "3", "1", "100","ALIVE"},
	}
	
	
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(3),
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

func (e *EmpireList) Update(msg tea.Msg) tea.Cmd {
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return changeMenu("ftm")
		case "enter":
			return tea.Batch(
				tea.Printf("Let's go to %s!", e.table.SelectedRow()[1]),
			)
		}
	}
	var cmd tea.Cmd
	e.table, cmd = e.table.Update(msg)
	return  cmd

}

func (e EmpireList) View() string {
	return baseStyle.Render(e.table.View()) + "\n"
}






