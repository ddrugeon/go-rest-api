package droids

import (
	"context"

	"github.com/ddrugeon/go-rest-api/internal/model"
	"github.com/ddrugeon/go-rest-api/internal/model/db"
)

// Service returns
type Service interface {
	Get(ctx context.Context) ([]model.Droid, error)
	GetByID(ctx context.Context, id string) (model.Droid, error)
}

//
type service struct {
	repo db.Repository
}

// NewService returns a new Service.
func NewService(r db.Repository) Service {
	return &service{repo: r}
}

func (s service) Get(ctx context.Context) ([]model.Droid, error) {
	droids := s.repo.Get(ctx)

	return droids, nil
}

func (s service) GetByID(ctx context.Context, id string) (model.Droid, error) {
	droid, err := s.repo.GetByID(ctx, id)

	return droid, err
}
