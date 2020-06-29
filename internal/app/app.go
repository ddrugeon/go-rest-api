package app

import (
	"github.com/ddrugeon/go-rest-api/internal/droids"
	"github.com/ddrugeon/go-rest-api/internal/healthz"
	"github.com/sirupsen/logrus"
)

// App represents current app implementation
type App struct {
	HealthzService healthz.Service
	DroidService   droids.Service
	Port           string
	Logger         *logrus.Logger
	db             droids.InMemory
	Version        Version
}

type Version struct {
	GitCommit string
	BuildDate string
	Version   string
}

func NewApp(logger *logrus.Logger, gitCommit string, buildDate string, version string) (*App, error) {
	return &App{
		HealthzService: healthz.NewService(),
		DroidService:   droids.NewService(droids.NewRepository()),
		Port:           "0.0.0.0:9090",
		Logger:         logger,
		Version: Version{
			GitCommit: gitCommit,
			BuildDate: buildDate,
			Version:   version,
		},
	}, nil
}
