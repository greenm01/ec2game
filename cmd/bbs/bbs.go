package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"flag"
	"errors"
)

const (
	exitFail = 1
)

func getPath() (string,error) {
       
	// Use FLAG to get command line paramenters
	pathPtr := flag.String("path", "", "path to door32.sys file")
	required := []string{"path"}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			return *pathPtr, errors.New("missing path to door32.sys directory")
		}
	}
	return *pathPtr, nil
        
}

func startClient() error {

    PORT := ":1992" 
    l, err := net.Listen("tcp", PORT)
    if err != nil {
            return err
    }
    defer l.Close()

    c, err := l.Accept()
    if err != nil {
            return err
    }

    for {
            netData, err := bufio.NewReader(c).ReadString('\n')
            if err != nil {
                    return err
            }
            if strings.TrimSpace(string(netData)) == "STOP" {
                    return err
            }

            fmt.Print("-> ", string(netData))
            t := time.Now()
            myTime := t.Format(time.RFC3339) + "\n"
            c.Write([]byte(myTime))
    }
	
	return nil
               
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}       
        
