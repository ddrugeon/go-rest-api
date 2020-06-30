package db

import (
	"context"

	"github.com/ddrugeon/go-rest-api/internal/model"
)

type Repository interface {
	Get(ctx context.Context) []model.Droid
	GetByID(ctx context.Context, id string) (model.Droid, error)
	Put(d model.Droid)
	Ping() (string, error)
}
