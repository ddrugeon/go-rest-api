package healthz

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
	return "OK", nil
}
