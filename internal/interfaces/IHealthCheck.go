package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckHadler interface {
	HealthChecHandlerHTTP(c *gin.Context)
}
type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepository interface {
}
