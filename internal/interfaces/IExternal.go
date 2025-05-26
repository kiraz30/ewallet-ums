package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternalService interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
}
