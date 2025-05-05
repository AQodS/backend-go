package middlewares

import (
	"backend-go/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// import SECRET_KEY from .env
var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("Authorization")

		// if token is empty, return 401 Unauthorized
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort() // abort the next request
			return
		}

		// Header format: Bearer <token>, delete "Bearer"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Create struct to token claim
		claims := &jwt.RegisteredClaims{}

		// Parse and verify token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		// if token is valid, set user_id to context
		c.Set("username", claims.Subject)

		// call next handler
		c.Next()
	}
}
