package droids

import (
	"context"
	"errors"
)

type Repository interface {
	Get(ctx context.Context) []Droid
	GetByID(ctx context.Context, id string) (*Droid, error)
	put(d Droid)
}

// InMemory is the in-memory concrete implementation of the repository.
type InMemory struct {
	droids map[string]Droid
}

func NewRepository() Repository {
	values := make(map[string]Droid)
	db := InMemory{droids: values}


	droid := Droid{
		ID:       "1",
		Name:     "R2-D2",
		Type:     "Droide astromecano",
		Company:  "Industrial Automation",
		Class:    "Droide Astromech",
		Model:    "Serie R2",
		Height:   "0,96",
		Vehicles: "X-Wing T65, Intercepteur Eta-2 classe Actis",
	}
	db.put(droid)

	droid = Droid{
		ID:       "2",
		Name:     "BB-8",
		Type:     "Droide astromecano",
		Class:    "Droide Astromech",
		Model:    "Unite BB",
		Height:   "0,67",
	}
	db.put(droid)

	droid = Droid{
		ID:       "3",
		Name:     "C-3PO",
		Company:  "Anakin Skywalker",
		Type:     "Droide Social",
		Class:    "Droide de protocole",
		Model:    "3PO",
		Height:   "1,67",
	}
	db.put(droid)

	return &db
}

// Get returns all droids.
func (im *InMemory) Get(ctx context.Context) []Droid {
	var values []Droid
	for _, value := range im.droids {
		values = append(values, value)
	} 
	return values
}

func (im *InMemory) GetByID(ctx context.Context, id string) (*Droid, error) {
	droid, found := im.droids[id]
	if !found {
		return nil, errors.New("No droid exists with ID: " + id)
	}

	return &droid, nil
}

func (im *InMemory) put(d Droid) {
	im.droids[d.ID] = d
}
