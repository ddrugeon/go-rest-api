package healthz

import (
	"math/rand"
)

// HealthStatus represents current health of microservice and it is used to serialize response.
type HealthStatus struct {
	Status string `json:"status"`
}

// Service returns
type Service interface {
	Health() HealthStatus
}

//
type service struct {
	status string
}

// NewService returns a new Service.
func NewService() Service {
	return &service{
		status: "0K",
	}
}

// Health returns current healths.
func (s *service) Health() HealthStatus {
	var status string
	if rand.Intn(10) > 3 {
		status = "Node ERROR: Node not responding"
	} else {
		status = "OK"
	}

	return HealthStatus{
		Status: status,
	}
}
