package bbsui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ti "github.com/charmbracelet/bubbles/textinput"
	//lg "github.com/charmbracelet/lipgloss"	

)

type Bio struct {
	textInput ti.Model
	valid bool
}

func (b Bio) GetText() string {
	return b.textInput.Value()	
}

func (b Bio) Init() tea.Cmd {
	return ti.Blink
}

func (b *Bio) Build() {
	
	b.textInput = ti.New()
	b.textInput.Placeholder = "Once upon a time..."
	b.textInput.Focus()
	b.textInput.CharLimit = 140
	b.textInput.Width = 70
	
	b.valid = false
	
}

func (b Bio) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	var cmd tea.Cmd

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyCtrlC, tea.KeyEsc:
					return arb.GetModel("ftm"), changeMenu("ftm")
				case tea.KeyCtrlS, tea.KeyEnter:
					b.valid = true
				case tea.KeyRunes:
					switch string(msg.Runes) {
						case "y", "Y":
						    if !b.valid { break }
							arb.Update("bio", b) 
							return arb.GetModel("ftm"), tea.Quit
						case "n", "N":
							if !b.valid { break }
							b.valid = false
							b.textInput.Reset()
							return b, cmd				
					}
			}
		case menuCmd:
			return b, ti.Blink	
	} 

	b.textInput, cmd = b.textInput.Update(msg)
	return b, cmd

}

func (b Bio) View() string {

	notice := "Tell us about your empire and place other players on notice:"	
	
	if !b.valid {
		return fmt.Sprintf(
				"%s\n\n%s\n\n%s",
				notice,
				b.textInput.View(),
				"(esc to quit)",
				) + "\n"
	} 

	return fmt.Sprintf(
			"%s\n\n> %s\n\n%s",
			notice,
			b.textInput.Value(),
			"Are you sure? [Y/N]",
			) + "\n"
}
