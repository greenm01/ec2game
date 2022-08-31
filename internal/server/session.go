package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Session struct {
	sid                       int
	gameSpace                 *GameSpace
	connection                net.Conn
	incoming                  chan string
	outgoing                  chan string
	reader                    *bufio.Reader
	writer                    *bufio.Writer
	killRoomConnGoroutine     chan bool
	killSocketReaderGoroutine chan bool
	killSocketWriterGoroutine chan bool
	sessionMutex              sync.Mutex
}

func NewSession(sid int, gameSpace *GameSpace, connection net.Conn) *Session {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	s := &Session{
		sid:                       sid,
		gameSpace:                 gameSpace,
		connection:                connection,
		incoming:                  make(chan string),
		outgoing:                  make(chan string),
		reader:                    reader,
		writer:                    writer,
		killRoomConnGoroutine:     make(chan bool),
		killSocketReaderGoroutine: make(chan bool),
		killSocketWriterGoroutine: make(chan bool),
		//		sessionMutex:
	}
	fmt.Println("A new session created. sid=", sid)
	return s
}

func (s *Session) Read() {
	for {
		select {
		case <-s.killSocketReaderGoroutine:
			return
		default:
			line, err := s.reader.ReadString('\n')
			if err != nil {
				// EOF? yes: disconnected
				// Judge. if true: LeaveAndDelete -- pop session from gameSpace, and delete session
				if err == io.EOF {
					fmt.Println("Client disconnected. Destroy session, sid=", s.sid)
					s.LeaveAndDelete()
				}

				// else:
				fmt.Println("bufio.reader.ReadString failed.")
				fmt.Println(err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			s.incoming <- line
		}
	}
}

func (s *Session) Write() {
	for {
		select {
		case <-s.killSocketWriterGoroutine:
			return
		case data := <-s.outgoing:
			s.writer.WriteString(data)
			s.writer.Flush()
		}
	}
}

func (s *Session) Listen() {
	go s.Read()
	go s.Write()
}

func (s *Session) LeaveAndDelete() {
	// leave

	gameSpace := *s.gameSpace
	sid := s.sid
	gameSpace.roomMutex.Lock()
	defer gameSpace.roomMutex.Unlock()
	delete(gameSpace.sessions, sid)

	// delete

	s.sessionMutex.Lock()
	defer s.sessionMutex.Unlock()

	// release resources

	// resouce: socket reader goroutine & socket writer goroutine
	s.killSocketReaderGoroutine <- true
	s.killSocketWriterGoroutine <- true

	// resource: reader & writer
	s.reader = nil
	s.writer = nil

	// resource: socket conection
	s.connection.Close()
	s.connection = nil

	// resource: RoomConnGoroutine
	s.killRoomConnGoroutine <- true

	// resource: connection to gameSpace
	// "many in, one out" 형태의 'gameSpace'의(!) channel 이기에 지워야할 채널(RoomConn)이 사실은 존재하지 않는다. (위에서) 해당 goroutine 지워줬으니 끝.
}
