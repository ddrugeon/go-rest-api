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

func NewApp(logger *logrus.Logger, gitCommit string, buildDate string, version string) (*App, error) {
	repo := db.NewRedisRepository(":6379")

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
		Port:           "0.0.0.0:9090",
		Logger:         logger,
		Version: Version{
			GitCommit: gitCommit,
			BuildDate: buildDate,
			Version:   version,
		},
		Repository: repo,
	}, nil
}
