package db

import (
	"context"
	"errors"

	"github.com/ddrugeon/go-rest-api/internal/model"
)

// InMemory is the in-memory concrete implementation of the repository.
type InMemory struct {
	droids map[string]model.Droid
}

func NewMemoryRepository() Repository {
	values := make(map[string]model.Droid)
	db := InMemory{droids: values}

	return &db
}

func (im *InMemory) Get(ctx context.Context) []model.Droid {
	var values []model.Droid
	for _, value := range im.droids {
		values = append(values, value)
	}
	return values
}

func (im *InMemory) GetByID(ctx context.Context, id string) (model.Droid, error) {
	droid, found := im.droids[id]
	if !found {
		return model.Droid{}, errors.New("No droid exists with ID: " + id)
	}

	return droid, nil
}

func (im *InMemory) Put(d model.Droid) error {
	im.droids[d.ID] = d

	return nil
}

func (im *InMemory) Ping() error {
	return nil
}

func (im *InMemory) Version() string {
	return "InMemory database"
}
