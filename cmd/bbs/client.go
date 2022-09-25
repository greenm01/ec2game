package main

import (
	"net"
	"bufio"
	"log"
	"strconv"
	//"strings"
	
	"github.com/pkg/errors"	
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/greenm01/ec2game/internal/bbsui"
	"github.com/greenm01/ec2game/internal/core"
)


// Open connects to a TCP Address.
// It returns a TCP connection armed with a timeout and wrapped into a
// buffered ReadWriter.
// https://github.com/AppliedGo/networking/blob/master/networking.go
func Open(addr string) (*bufio.ReadWriter, error) {
	// Dial the remote process.
	// Note that the local port is chosen on the fly. If the local port
	// must be a specific one, use DialTCP() instead.
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

type bbsClient struct {
	user string
	ip string
	port string
	tui *tea.Program
	buf bufio.ReadWriter
	state core.PlayerState
}

func (c *bbsClient) Run(m ui.Menu) error {
	
	if err := c.start(); err != nil {
		return err
	}
	
	ui.Cls()
	c.tui = tea.NewProgram(m)
	if err := c.tui.Start(); err != nil {
		return err
	}
	ui.Cls()
	return nil
}

func (c *bbsClient) start() error {

	addr := c.ip + ":" + c.port
		
	// Open a connection to the server.
	rw, err := Open(addr)
	if err != nil {
		return errors.Wrap(err, "Client: Failed to open connection to "+addr)
	}	
	
	// Send a USER request.
	log.Println("Send the USER request.")
	n, err := rw.WriteString("USER: " + c.user + "\n")
	if err != nil {
		return errors.Wrap(err, "Could not send the USER request ("+strconv.Itoa(n)+" bytes written)")
	}	
	
	if err = flushBuffer(rw); err != nil { return err }

	// Read the reply.
	log.Println("Read the reply.")
	c.state, err = core.DecodePlayerState(rw)
	if err != nil {
		return errors.Wrap(err, "Client: Failed to read the reply.")
	}	
	log.Println("USER request: got " + c.state.User.Name + "'s game data.")

	if err = flushBuffer(rw); err != nil { return err }
	
	for {
		
	}
	
	return nil
               
}

func flushBuffer(rw *bufio.ReadWriter) error {
	log.Println("Flush the buffer.")
	if err := rw.Flush(); err != nil {
		return errors.Wrap(err, "Flush failed.")
	}
	return nil
}