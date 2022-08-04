package ui

import (
	"github.com/rivo/tview"
	"log"
	"time"
)

const (
	offsetRange = 100
	loadDelay   = time.Millisecond * 50
	maxOffset   = 10000
)

// MainPage : This struct contains the grid and the entry table.
// In addition, it also keeps track of whether to show followed/popular manga based on login status
// as well as the entry offset.
type MainPage struct {
	Grid          *tview.Grid  // The page grid.
	Table         *tview.Table // The table that contains the list of manga.
	CurrentOffset int
	MaxOffset     int

	//cWrap *utils.ContextWrapper // For context cancellation.
}

// ShowMainPage : Make the app show the main page.
func ShowMainPage() {
	// Create the new main page
	log.Println("Creating new main page...")
	//mainPage := newMainPage()

	//core.App.TView.SetFocus(mainPage.Grid)
	//core.App.PageHolder.AddAndSwitchToPage(utils.MainPageID, mainPage.Grid, true)
}