package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSVC := &services.HealthCheck{}
	healtCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSVC,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerSVC := &services.RegisterService{
		RegisterRepo: registerRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSVC,
	}

	r := gin.Default()
	r.GET("/health", healtCheckAPI.HealthChecHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerAPI.Register)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
