package main

import (
	"fmt"
	"os/exec"
)

func runServer() {
	// Create command.
	// ... The application is not executed yet.
	cmd := exec.Command("server.exe")
	fmt.Println("Starting command")
	// Run firefox-bin.
	// ... This starts a web browser instance.
	//     It does not wait for it to finish.
	cmd.Start()
	fmt.Println("DONE")
}
