//go:build mage

package main

import (
	"fmt"
	"runtime"

	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/mg"	
)

func checkOS () string {
	if runtime.GOOS == "windows" {
		return "ec2game.exe"
	} else {
		return "ec2game"
	}	
}

func Build() error {
	fmt.Println("Build running")
	mg.Deps(Clean)
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	
	if err := sh.Run("go", "build","-ldflags", "-s -w", "./cmd/ec2game"); err != nil {
		return err
	}

	fmt.Println("compressing executible with upx")	
	return sh.Run("upx","-9",checkOS())
}

// Remove the temporarily generated files from Release.
func Clean() error {
	return sh.Rm(checkOS())
}


