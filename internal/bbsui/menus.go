package bbsui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func Cls() {
	fmt.Print("\033[H\033[2J")	
}

// address keeps track of the current menu
var address arbiter

type arbiter struct {
	address string
}

func (a *arbiter) Get() string {
	return a.address
}

func (a *arbiter) Set(adr string) {
	a.address = adr
}

type menuCmd bool

func changeMenu(adr string) tea.Cmd {
	address.Set(adr)
	var m menuCmd
	return func() tea.Msg {
		Cls()
		return m
	} 
}

type Model interface {
	Update(tea.Msg) (tea.Cmd)
	View() string
}

func UDate (m Model, msg tea.Msg) tea.Cmd  {
	return m.Update(msg)
}

type Menu struct {
	MenuMap map[string]Model
}

func (m *Menu) Build(address string, model Model) {
	if address == "" {
		panic("address should not be empty")
	}
	if model == nil {
		panic("model should not be nil")
	}
	if m.MenuMap == nil {
		m.MenuMap = make(map[string]Model)
	}
	m.MenuMap[address] = model
}

func (m Menu) Init() tea.Cmd {
	return changeMenu("ftm")
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	adr := address.Get()
	cmd := UDate(m.MenuMap[adr],msg)
	return m,cmd
}

func (m Menu) View() string {
	adr := address.Get()
	return m.MenuMap[adr].View()
}

/*
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
	
}*/