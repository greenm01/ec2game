package bbsui

import (
	"fmt"
	"github.com/greenm01/ec2game/internal/core"
	tea "github.com/charmbracelet/bubbletea"
)

func Cls() {
	fmt.Print("\033[H\033[2J")	
}

// address keeps track of the current menu
var arb arbiter

type arbiter struct {
	modelMap map[string]tea.Model
	address string
}

func (a *arbiter) Get() string {
	return a.address
}

func (a *arbiter) Set(adr string) {
	a.address = adr
}

func (a *arbiter) GetModel(adr string) tea.Model {
	if adr == "bio" {
		a.modelMap["bio"],_ = a.modelMap["bio"].Update(tea.KeyEnter)
	}
	return a.modelMap[adr]
}

func (a *arbiter) Update(adr string, m tea.Model) {
	a.modelMap[adr] = m
	a.address = adr
}

func (a *arbiter) Add(address string, model tea.Model) {
	if address == "" {
		panic("address should not be empty")
	}
	if model == nil {
		panic("model should not be nil")
	}
	if a.modelMap == nil {
		a.modelMap = make(map[string]tea.Model)
	}
	a.modelMap[address] = model
}

type menuCmd bool

func changeMenu(adr string) tea.Cmd {
	arb.Set(adr)
	var m menuCmd
	m = true
	return func() tea.Msg {
		Cls()
		return m
	} 
}

// ftm initializes the First Time Menu
func FtmSetup(ps core.PlayerState) tea.Model {
	
	ftm := FirstTime{}
	ftm.Build()
	arb.Add("ftm", ftm)
	
	pager := Pager{}
	pager.Build(Intro())
	arb.Add("intro", pager)
	
	el := EmpireList{}
	el.Build(ps)
	arb.Add("empires", el)
	
	join := Join{}
	join.Build()
	arb.Add("join", join)
	
	bio := Bio{}
	bio.Build()
	arb.Add("bio", bio)
	
	return ftm
	
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