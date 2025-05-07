package main

import (
	_ "mtb_web/cmd/server/docs"
	"mtb_web/internal/initialize"
)

func main() {
	// This is the entry point for the server application.
	// The server will be started here.
	initialize.Run()
}
