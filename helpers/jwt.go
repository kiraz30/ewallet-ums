package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"full_name"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

var jwtSecret = []byte(GetEnv("APP_SECRET", ""))

func GenerateToken(ctx context.Context, userID int, username, fullname, tokenType string, now time.Time) (string, error) {
	now = time.Now()
	claimToken := ClaimToken{
		UserID:   userID,
		Username: username,
		Fullname: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
		},
	}
	log.Println("Generated ExpiresAt:", now.Add(MapTypeToken[tokenType]))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)
	resulToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return resulToken, fmt.Errorf("failed to generate token: %v", err)
	}
	return resulToken, nil

}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {
	var (
		claimToken *ClaimToken
		ok         bool
	)
	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	if claimToken, ok = jwtToken.Claims.(*ClaimToken); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claimToken, nil
}
