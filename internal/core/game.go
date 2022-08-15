package core

import (
	"fmt"
	"os"

	"github.com/greenm01/ec2game/internal/io"

	tea "github.com/charmbracelet/bubbletea"
)

// AppVersion : Global App Version
const AppVersion = "EC2 v0.1"

// Root game model
type EC2 struct {
	frame *io.GameFrame
}

func (g EC2) runGame() error {

	p := tea.NewProgram(g.frame)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil

}

func InitGame() error {

	game := EC2{frame: new(io.GameFrame)}
	game.frame.InitCmd()
	game.runGame()

	return nil
}
