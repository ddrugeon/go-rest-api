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

	droid := model.Droid{
		ID:       "1",
		Name:     "R2-D2",
		Type:     "Droide astromecano",
		Company:  "Industrial Automation",
		Class:    "Droide Astromech",
		Model:    "Serie R2",
		Height:   "0,96",
		Vehicles: "X-Wing T65, Intercepteur Eta-2 classe Actis",
	}
	db.Put(droid)

	droid = model.Droid{
		ID:     "2",
		Name:   "BB-8",
		Type:   "Droide astromecano",
		Class:  "Droide Astromech",
		Model:  "Unite BB",
		Height: "0,67",
	}
	db.Put(droid)

	droid = model.Droid{
		ID:      "3",
		Name:    "C-3PO",
		Company: "Anakin Skywalker",
		Type:    "Droide Social",
		Class:   "Droide de protocole",
		Model:   "3PO",
		Height:  "1,67",
	}
	db.Put(droid)

	return &db
}

// Get returns all droids.
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

func (im *InMemory) Put(d model.Droid) {
	im.droids[d.ID] = d
}

func (im *InMemory) Ping() (string, error) {
	return "OK", nil
}
