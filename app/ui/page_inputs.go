package ui

import (
	"github.com/greenm01/ec2game/app/core"
	"github.com/gdamore/tcell/v2"
	//"github.com/rivo/tview"	
)

// SetUniversalHandlers : Set universal inputs for the app.
func SetUniversalHandlers() {
	// Set universal keybindings
	core.App.TView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlL: // Login/Logout
			ctrlLInput()
		case tcell.KeyCtrlK: // Help page.
			ctrlKInput()
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

func ctrlKInput() {
	
}