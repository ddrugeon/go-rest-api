package main

import (
	"fmt"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/server"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Println("Got Error during initialization")
	}

	server := server.NewServer(app)
	server.Run(":9090")
}
