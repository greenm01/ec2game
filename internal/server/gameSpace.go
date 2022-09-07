package server

import (
	"fmt"
	"net"
	"sync"
	
	"github.com/greenm01/ec2game/internal/core"
	
)

type GameSpace struct {
	
	// Network handlers
	sessions  map[int]*Session
	lastSid   int
	entrance  chan net.Conn
	incoming  chan string
	outgoing  chan string
	roomMutex sync.Mutex
	
	// Game related
	cfg core.Config
	gState core.GameState
	
}

func NewGameSpace(cg core.Config, gs core.GameState) *GameSpace {
	space := &GameSpace{
		sessions: make(map[int]*Session),
		lastSid:  -1,
		entrance: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
		cfg: cg,
		gState: gs,
	}
	fmt.Println("A new GameSpace created.")
	return space
}

func (gs *GameSpace) Connect(conn net.Conn) {
	gs.entrance <- conn
	return
}

func (gs *GameSpace) Broadcast(data string) {
	for _, session := range gs.sessions {
		session.outgoing <- data
	}
}

func (gs *GameSpace) Join(connection net.Conn) {
	gs.roomMutex.Lock()
	defer gs.roomMutex.Unlock()
	newSessionId := gs.lastSid + 1
	gs.lastSid = newSessionId

	session := NewSession(newSessionId, gs, connection)
	session.Listen()
	fmt.Println("session started listening.")

	_, keyExist := gs.sessions[newSessionId]
	if !keyExist {
		gs.sessions[newSessionId] = session
	}

	go func() { // goroutine for roomConn writer
		for {
			select {
			case <-session.killRoomConnGoroutine:
				return
			case data := <-session.incoming:
				gs.incoming <- data
			}
		}
	}()
}

// This goroutine runs forever, and does not need a channel to kill it.
func (gs *GameSpace) Listen() {
	go func() {
		for {
			select {
			case data := <-gs.incoming:
				//fmt.Println("RECEIVED: " + data)
				gs.Broadcast(data)
			case conn := <-gs.entrance:
				gs.Join(conn)
			}
		}
	}()
}
