package main

import (
	"backend-go/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Gin initialization
	router := gin.Default()

	// Create route with GET method
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	// Run the server
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
