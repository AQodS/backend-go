package helpers

import (
	"backend-go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {
	// expired token in 60 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	// create claims JWT, set subject and expiration time
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// create new token with claims and use HS256 algorithm
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return token
}
