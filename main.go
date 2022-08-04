// EC2 game

package main

import "github.com/greenm01/ec2game/app/service"

func main() {
	// Initialize the app
	service.Start()
	defer service.Shutdown()
}

