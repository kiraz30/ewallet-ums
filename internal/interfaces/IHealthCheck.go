package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepository interface {
}

type IHealthCheckHandler interface {
	HealthChecHandlerHTTP(c *gin.Context)
}
