package main

import (
    "time"
	"github.com/greenm01/ec2game/internal/core"
)

type gameState struct {
    Started bool
    LaunchTime time.Time
    StarMap *core.StarMap
    Empires map[int]*core.Empire
}
