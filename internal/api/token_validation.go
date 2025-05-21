package api

import (
	"context"
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/constans"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"fmt"
)

type TokenValidationHandler struct {
	tokenvalidation.UnimplementedTokenValidationServer
	TokenValidationService interfaces.ITokenValidationService
}

// nama harus disamakan dengan nama validateToken di file grpc.pb.go
func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("Token is empty")
		log.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}
	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constans.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
			Email:    claimToken.Email,
		},
	}, nil
}
