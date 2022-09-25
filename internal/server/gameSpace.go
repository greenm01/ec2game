package server

import (
	"fmt"
	"net"
	"sync"
	"strings"
	"github.com/greenm01/ec2game/internal/core"
	"github.com/greenm01/ec2game/internal/util"
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
	state core.GameState
	
}

func NewGameSpace(cg core.Config, gs core.GameState) *GameSpace {
	space := &GameSpace{
		sessions: make(map[int]*Session),
		lastSid:  -1,
		entrance: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
		cfg: cg,
		state: gs,
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

	go func() { 
		for {
			select {
			case <-session.killRoomConnGoroutine:
				return
			case data := <-session.incoming:
				
				//gs.incoming <- data
				/* TODO: Lookup user, load playerstate, broadcast to client  */	
				s := util.Substr(data,0,5)
				
				if s == "USER:" {
					// user login
					name := strings.TrimSpace(util.Substr(data,6,len(data)))
					
					var ps core.PlayerState
					var user *core.User
					var found bool
			
					if user, found = gs.state.Users[name]; !found {
						// User not in game
						user = &core.User{ID:-1, Name:name, FirstTime:true}
					}
					 
					ps = core.PlayerState{User:user}					
					ps.Setup(gs.state)
					d,_ := ps.Encode()
					session.outgoing <- d.String()

					fmt.Println("Hello,",name)
				
				} else if s == "JOIN:" {
					// new user join game
					
				} else {
					// process commands
				}				
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
				gs.Broadcast(data)
			case conn := <-gs.entrance:
				gs.Join(conn)
			}
		}
	}()
}
