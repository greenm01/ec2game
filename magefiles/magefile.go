//go:build mage

package main

import (
	"fmt"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func checkOS() string {
	if runtime.GOOS == "windows" {
		return "ec2game.exe"
	} else {
		return "ec2game"
	}
}

func Build() error {
	fmt.Println("Building executible...")
	mg.Deps(Clean)

	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	return sh.Run("go", "build", "-ldflags", "-s -w", "./cmd/ec2game")
}

// Remove the temporarily generated files from Release.
func Clean() error {
	return sh.Rm(checkOS())
}
