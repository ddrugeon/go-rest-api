package main

import (
	"fmt"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/server"
)

func main() {
	const addr = "0.0.0.0:9090"
	app, err := app.NewApp()
	if err != nil {
		fmt.Println("Got Error during initialization")
	}

	server := server.NewServer(app)
	server.Run(addr)
}
