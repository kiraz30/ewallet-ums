package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request models.LoginRequest) (models.LoginReponse, error) {

	var (
		response models.LoginReponse
		now      = time.Now().UTC()
	)
	userDetail, err := s.UserRepo.GetUserByUserName(ctx, request.Username)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user by username")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}
	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	//session
	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert new user session")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}
