//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/mg"	
)

const (
	ec2loc = "./cmd/ec2/"
	loc2 = "./..."
)


func Build() error {
	fmt.Println("Build running")
	mg.Deps(Clean)
	
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", ec2loc)
}

// Remove the temporarily generated files from Release.
func Clean() error {
	return sh.Rm("ec2")
}


