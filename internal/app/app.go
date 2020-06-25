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
}

func NewApp(logger *logrus.Logger) (*App, error) {
	return &App{
		HealthzService: healthz.NewService(),
		DroidService:   droids.NewService(droids.NewRepository()),
		Port:           "localhost:9090",
		Logger:         logger,
	}, nil
}
