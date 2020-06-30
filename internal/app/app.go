package app

import (
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
	repo := db.NewMemoryRepository()

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
	}, nil
}
