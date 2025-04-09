package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IRegisterService interface {
	Regiter(ctx context.Context, request models.User) (interface{}, error)
}
