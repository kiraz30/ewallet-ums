package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type ILoginService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginReponse, error)
}
