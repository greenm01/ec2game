package bbsui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ti "github.com/charmbracelet/bubbles/textinput"
	//lg "github.com/charmbracelet/lipgloss"	

)

type Join struct {
	textInput ti.Model
	valid bool
}

func (j Join) Init() tea.Cmd {
	return ti.Blink
}

func (j *Join) Build() {
	
	j.textInput = ti.New()
	j.textInput.Placeholder = "Esterian Conquest"
	j.textInput.Focus()
	j.textInput.CharLimit = 25
	j.textInput.Width = 25
	
	j.valid = false
	
}

func (j Join) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	var cmd tea.Cmd

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyCtrlC, tea.KeyEsc:
					return arb.GetModel("ftm"), changeMenu("ftm")
				case tea.KeyCtrlS, tea.KeyEnter:
					j.valid = true
				case tea.KeyRunes:
					switch string(msg.Runes) {
						case "y", "Y":
						    if !j.valid { break } 
							return arb.GetModel("bio"), changeMenu("bio")
						case "n", "N":
							if !j.valid { break }
							j.valid = false 
							j.textInput.Reset()
							return j, cmd				
					}
			}
		case menuCmd:
			return j, ti.Blink
	}	

	j.textInput, cmd = j.textInput.Update(msg)
	return j, cmd

}

func (j Join) View() string {
	
	if !j.valid {
		return fmt.Sprintf(
				"Name your empire:\n\n%s\n\n%s",
				j.textInput.View(),
				"(esc to quit)",
				) + "\n"
	} 

	return fmt.Sprintf(
			"Name your empire:\n\n> %s\n\n%s",
			j.textInput.Value(),
			"Are you sure? [Y/N]",
			) + "\n"
}

