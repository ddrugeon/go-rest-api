package main

import (
	"log"
	"os"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/server"
	"github.com/sirupsen/logrus"
)

var (
	// GitCommit holds short commit hash of source tree
	GitCommit string

	// GitBranch holds current branch name the code is built off
	GitBranch string

	// GitState shows whether there are uncommitted changes
	GitState string

	// GitSummary holds output of git describe --tags --dirty --always
	GitSummary string

	// BuildDate holds RFC3339 formatted UTC date (build time)
	BuildDate string

	// Version holds contents of ./VERSION file, if exists, or the value passed via the -version option
	Version string
)

func main() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.SetFormatter(&logrus.JSONFormatter{})

	databaseUrl := os.Getenv("REDIS_URL")
	port := os.Getenv("PORT")

	logger.Debugln("GitCommit: ", GitCommit)
	logger.Debugln("BuildDate: ", BuildDate)
	logger.Debugln("Version: ", Version)

	app, err := app.NewApp(logger, GitCommit, BuildDate, Version, databaseUrl, port)
	if err != nil {
		log.Println("Got Error during initialization")
	}

	server := server.NewServer(app)
	server.Run()
}
