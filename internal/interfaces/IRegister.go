package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Regiter(ctx context.Context, request models.User) (interface{}, error)
}
type IRegisterHandler interface {
	Register(c *gin.Context)
}
