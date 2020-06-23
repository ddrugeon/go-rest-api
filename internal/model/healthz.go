package model

type HealthStatus struct {
	Status string `json:"status"`
}

func NewHealthStatus(status string) *HealthStatus {
	return &HealthStatus{
		Status: status,
	}
}
