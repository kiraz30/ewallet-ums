package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterHadler interface {
	Register(c *gin.Context)
}

type IRegisterService interface {
	Regiter(ctx context.Context, request models.User) (interface{}, error)
}
