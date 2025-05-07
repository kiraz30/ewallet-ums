package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.GET("/health", dependency.HealthCheckApi.HealthChecHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterApi.Register)
	userV1.POST("/login", dependency.LoginApi.Login)

	userV1WithAuth := userV1.Use()
	userV1WithAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutApi.Logout)
	userV1WithAuth.PUT("/refresh-token", dependency.MiddlewareValidateRefreshToken, dependency.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository  interfaces.IUserRepository
	HealthCheckApi  interfaces.IHealthCheckHandler
	RegisterApi     interfaces.IRegisterHandler
	LoginApi        interfaces.ILoginHandler
	LogoutApi       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler
}

func dependencyInject() Dependency {
	healthCheckSVC := &services.HealthCheck{}
	healtCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSVC,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	registerSVC := &services.RegisterService{
		UserRepo: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSVC,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}

	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}

	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}
	RefreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}
	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: RefreshTokenSvc,
	}
	return Dependency{
		HealthCheckApi:  healtCheckAPI,
		RegisterApi:     registerAPI,
		LoginApi:        loginAPI,
		LogoutApi:       logoutAPI,
		UserRepository:  userRepo,
		RefreshTokenAPI: refreshTokenAPI,
	}

}
