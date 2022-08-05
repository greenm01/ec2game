package ui

import (
	"github.com/greenm01/ec2game/internal/core"
	"github.com/greenm01/ec2game/internal/ui/utils"
	"github.com/gdamore/tcell/v2"
	//"github.com/rivo/tview"	
)

// SetUniversalHandlers : Set universal inputs for the app.
func SetUniversalHandlers(g *core.EC2) {
	// Set universal keybindings
	g.TView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlL: // Login/Logout
			ctrlLInput()
		case tcell.KeyCtrlK: // Help page.
			ctrlKInput(g)
		case tcell.KeyCtrlS: // Search page.
			ctrlSInput()
		case tcell.KeyCtrlC: // Ctrl-C interrupt.
			ctrlCInput()
		}
		return event // Forward the event to the actual current primitive.
	})
}

func ctrlCInput() {
	
}

func ctrlSInput() {
	
}

func ctrlLInput() {
	
}

func ctrlKInput(g *core.EC2) {
	ShowHelpPage(g)
}

// setHandlers : Set handlers for the help page.
func (p *HelpPage) setHandlers(g *core.EC2) {
	// Set grid input captures.
	p.Grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			g.PageHolder.RemovePage(utils.HelpPageID)
		}
		return event
	})
}
