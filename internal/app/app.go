package app

import (
	"github.com/ddrugeon/go-rest-api/internal/healthz"
)

// App represents current app implementation
type App struct {
	HealthzService healthz.Service
}

func NewApp() (*App, error) {
	return &App{
		HealthzService: healthz.NewService(),
	}, nil
}
