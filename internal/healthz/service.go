package healthz

import (
	"errors"
	"math/rand"
)

// Service returns
type Service interface {
	Health() (string, error)
}

//
type service struct {
}

// NewService returns a new Service.
func NewService() Service {
	return &service{}
}

// Health returns current healths.
func (s *service) Health() (string, error) {
	var result string
	var err error
	if rand.Intn(10) > 3 {
		err = errors.New("Node ERROR: Node not responding")
		result = ""
	} else {
		err = nil
		result = "OK"
	}

	return result, err
}
