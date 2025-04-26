package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type ILoginHandler interface {
	Login(c *gin.Context)
}

type ILoginService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginReponse, error)
}
