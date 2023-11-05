package main

import (
	cli "notes/app/cli"
	"os"
)

// @title Nøtes API
// @version 0.0.1
// @description A simple REST API for a simple note taking app.

// @license.name MIT
// @license.url https://github.com/TheBigRoomXXL/Notes/raw/main/LICENCE.md

// @contact.name Github
// @contact.url https://github.com/theBigRoomXXL/notes/
func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
