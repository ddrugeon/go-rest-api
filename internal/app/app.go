package app

import (
	"github.com/ddrugeon/go-rest-api/internal/model"
	"github.com/ddrugeon/go-rest-api/internal/model/db"
	"github.com/ddrugeon/go-rest-api/internal/services/droids"
	"github.com/ddrugeon/go-rest-api/internal/services/healthz"
	"github.com/sirupsen/logrus"
)

// App represents current app implementation
type App struct {
	HealthzService healthz.Service
	DroidService   droids.Service
	Port           string
	Logger         *logrus.Logger
	Repository     db.Repository
	Version        Version
}

type Version struct {
	GitCommit string
	BuildDate string
	Version   string
}

func NewApp(logger *logrus.Logger, gitCommit string, buildDate string, version string, database_url string, server_port string) (*App, error) {
	if database_url == "" {
		database_url = "localhost"
	}

	if server_port == "" {
		server_port = "9090"
	}

	logger.Debugln("Database URL: ", database_url)
	logger.Debugln("Server port: ", server_port)

	repo := db.NewRedisRepository(database_url + ":6379")

	droid := model.Droid{
		ID:       "droid:1",
		Name:     "R2-D2",
		Type:     "Droide astromecano",
		Company:  "Industrial Automation",
		Class:    "Droide Astromech",
		Model:    "Serie R2",
		Height:   "0,96",
		Vehicles: "X-Wing T65, Intercepteur Eta-2 classe Actis",
	}
	repo.Put(droid)

	droid = model.Droid{
		ID:     "droid:2",
		Name:   "BB-8",
		Type:   "Droide astromecano",
		Class:  "Droide Astromech",
		Model:  "Unite BB",
		Height: "0,67",
	}
	repo.Put(droid)

	droid = model.Droid{
		ID:      "droid:3",
		Name:    "C-3PO",
		Company: "Anakin Skywalker",
		Type:    "Droide Social",
		Class:   "Droide de protocole",
		Model:   "3PO",
		Height:  "1,67",
	}
	repo.Put(droid)

	return &App{
		HealthzService: healthz.NewService(repo),
		DroidService:   droids.NewService(repo),
		Port:           "0.0.0.0:" + server_port,
		Logger:         logger,
		Version: Version{
			GitCommit: gitCommit,
			BuildDate: buildDate,
			Version:   version,
		},
		Repository: repo,
	}, nil
}
