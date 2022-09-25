package main

import (
	"fmt"
	"os"
	"flag"
	"errors"
)

const (
	exitFail = 1
)

/* TODO: Make version a global constant */
func usage() string {
	return 	"Esterian Conquest v2.0\n\n" +
		 	"Invalid command line flags!\n\n" +
		   	"The required flags are:\n\n" +
	       	"          -drop      Path to door32.sys dropfile\n" +
	       	"          -game      Path to game diretory\n\n" +
	       	"Example: -drop ./bbsDropFilePath -game ./game1Path\n" 
}

func getPaths(dp string, gp string) (string,string,error) {
       
	// Use FLAG to get command line paramenters
	pathPtr := flag.String(dp, "", "path to door32.sys dropfile") 
	gamePth := flag.String(gp, "", "path to game files")
	required := []string{dp,gp}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			return *pathPtr, *gamePth, errors.New(usage())
		}
	}
	return *pathPtr, *gamePth, nil
	        
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}       
        
