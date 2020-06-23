package app

import (
	"github.com/ddrugeon/go-rest-api/internal/droids"
	"github.com/ddrugeon/go-rest-api/internal/healthz"
)

// App represents current app implementation
type App struct {
	HealthzService healthz.Service
	DroidService   droids.Service
	db             droids.InMemory
}

func NewApp() (*App, error) {
	return &App{
		HealthzService: healthz.NewService(),
		DroidService:   droids.NewService(droids.NewRepository()),
	}, nil
}
