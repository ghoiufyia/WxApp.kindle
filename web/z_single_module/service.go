package modules

import ()

type Service struct {
	Name	string
	age		int32
}

// NewService returns a new Service instance
func NewService(name string,age int32) *Service {
	return &Service{Name:name,age:age}
}

// Close stops any running services
func (s *Service) Close() {}