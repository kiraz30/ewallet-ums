package interfaces

type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepository interface {
}
