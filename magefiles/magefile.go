//go:build mage

package main

import (
	"fmt"
	"runtime"
	"errors"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

func checkWin() bool {
	if runtime.GOOS == "windows" { return true }
	return false
}

// Builds the client app
func (Build) Game() error {
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	
	fmt.Print("Building game client...")
	
	if err := sh.Run("go", "build", "-ldflags", "-s -w", "./cmd/ec2g"); err != nil {
		return err
	}
	
	fmt.Println("success!")
	return nil
	
}

// Builds the game server
func (Build) Server() error {
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	
	fmt.Print("Building server...")
	
	if err := sh.Run("go", "build", "-ldflags", "-s -w", "./cmd/ec2s"); err != nil {
		return err
	}
	
	fmt.Println("success!")
	return nil
}

// Builds everything
func (Build) All() error {
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	
	fmt.Print("Building everything...")
	
	if err := sh.Run("go", "build", "-ldflags", "-s -w", "./cmd/ec2s"); err != nil {
		return err
	}
	
	if err := sh.Run("go", "build", "-ldflags", "-s -w", "./cmd/ec2g"); err != nil {
		return err
	}
	
	fmt.Println("success!")
	return nil
	
}

// Remove the temporarily generated files from Release.
func Clean() error {
	fmt.Print("Cleaning house....")
	if checkWin() == true {
		if sh.Rm("ec2g.exe") != nil || sh.Rm("ec2s.exe") != nil {
			return errors.New("remove file error...")
		}
		fmt.Println("succes!")
		return nil
	}  
	if sh.Rm("ec2g") != nil || sh.Rm("ec2s") != nil { 
		return errors.New("remove file error...")
	}
	fmt.Println("success!")
	return nil
}
