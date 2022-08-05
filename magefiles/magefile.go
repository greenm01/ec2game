//go:build mage

package main

import (
	"fmt"
	"runtime"

	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/mg"	
)

func Build() error {
	fmt.Println("Build running")
	mg.Deps(Clean)
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build","./cmd/ec2game")
}

// Remove the temporarily generated files from Release.
func Clean() error {
	if runtime.GOOS == "windows" {
		fmt.Println("Detected Windows OS")
		return sh.Rm("ec2game.exe")
	} else {
		fmt.Println("Detected POSIX")
		return sh.Rm("ec2game")
	}
}


