package main

import (
	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Level = logrus.InfoLevel
	log.SetFormatter(&logrus.JSONFormatter{})

	app, err := app.NewApp(log)
	if err != nil {
		log.Println("Got Error during initialization")
	}

	server := server.NewServer(app)
	server.Run()
}
