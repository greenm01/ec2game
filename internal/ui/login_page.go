package ui

import (
	//"github.com/greenm01/ec2game/internal/core"
	//"github.com/greenm01/ec2game/app/ui/utils"
	"github.com/rivo/tview"
	//"log"
)

// LoginPage : This struct contains the grid and form for the login page.
type LoginPage struct {
	Grid *tview.Grid
	Form *tview.Form
}

// ShowLoginPage : Make the app show the login page.
func ShowLoginPage() {
	// Create the new login page
	//loginPage := newLoginPage()

	//core.App.TView.SetFocus(loginPage.Grid)
	//core.App.PageHolder.AddAndSwitchToPage(utils.LoginPageID, loginPage.Grid, true)
}

// newLoginPage : Creates a new login page.
func newLoginPage() *LoginPage {
	// Create the LoginPage
	loginPage := &LoginPage{}
	
	return loginPage
}
