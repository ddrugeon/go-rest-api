package droids

import "context"

// Service returns
type Service interface {
	Get(ctx context.Context) ([]Droid, error)
	GetByID(ctx context.Context, id string) (*Droid, error)
}

//
type service struct {
	repo Repository
}

// NewService returns a new Service.
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s service) Get(ctx context.Context) ([]Droid, error) {
	droids := s.repo.Get(ctx)

	return droids, nil
}

func (s service) GetByID(ctx context.Context, id string) (*Droid, error) {
	droid, err := s.repo.GetByID(ctx, id)

	return droid, err
}
