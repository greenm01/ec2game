package main

import (
	"github.com/greenm01/ec2game/internal/core"
)

type gameState struct {
    Started bool
    Config *configData
    StarMap *core.StarMap
    Empires map[int]*core.Empire
}
