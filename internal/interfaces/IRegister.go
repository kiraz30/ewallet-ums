package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}
type IRegisterService interface {
	Regiter(ctx context.Context, request models.User) (interface{}, error)
}
