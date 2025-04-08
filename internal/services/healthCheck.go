package services

import "ewallet-framework/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service healty", nil
}
