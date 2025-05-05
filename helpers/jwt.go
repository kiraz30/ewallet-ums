package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

func GenerateToken(ctx context.Context, userID int, username, fullname, tokenType string, now time.Time) (string, error) {
	jwtSecret := []byte(GetEnv("APP_SECRET", ""))

	claim := ClaimToken{
		Username: username,
		Fullname: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	resulToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return resulToken, fmt.Errorf("failed to generate token: %v", err)
	}
	return resulToken, nil

}

func ValidateToken(ctx context.Context, tokenString string) (*ClaimToken, error) {
	jwtSecret := []byte(GetEnv("APP_SECRET", ""))

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{"HS256"}), // hanya izinkan HS256
		jwt.WithoutClaimsValidation(),           // <<== ini penting
	)
	// Parse token
	token, err := parser.ParseWithClaims(tokenString, &ClaimToken{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}

	// Ambil claim jika token valid
	claims, ok := token.Claims.(*ClaimToken)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Tambahkan log ini
	fmt.Println("Now:", time.Now())
	fmt.Println("Token expires at:", claims.ExpiresAt.Time)

	return claims, nil
}
