package service

import (
	"github.com/rivo/tview"
	"log"

	"github.com/greenm01/ec2game/app/core"
	"github.com/greenm01/ec2game/app/ui"
)

// Start : Set up the application.
func Start() {
	// Create new app.
	core.App = &core.EC2 {
		//Client:     mangodex.NewDexClient(),
		TView:      tview.NewApplication(),
		PageHolder: tview.NewPages(),
	}

	// Show appropriate screen based on restore session result.
	if err := core.App.Initialise(); err != nil {
		ui.ShowLoginPage()
	} else {
		ui.ShowMainPage()
	}
	log.Println("Initialised starting screen.")
	ui.SetUniversalHandlers()

	// Run the app.
	log.Println("Running app...")
	if err := core.App.TView.Run(); err != nil {
		log.Println(err)
	}
}

// Shutdown : Shutdown the application.
func Shutdown() {
	core.App.Shutdown()
}