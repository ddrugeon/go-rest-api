package healthz

import "github.com/ddrugeon/go-rest-api/internal/model/db"

// Service returns
type Service interface {
	Health() (string, error)
}

//
type service struct {
	db db.Repository
}

// NewService returns a new Service.
func NewService(r db.Repository) Service {
	return &service{
		db: r,
	}
}

// Health returns current healths.
func (s *service) Health() (string, error) {
	_, err := s.db.Ping()

	if err != nil {
		return "NOK", err
	}

	return "OK", nil
}
