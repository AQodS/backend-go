package main

import (
	"backend-go/config"
	"backend-go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	database.InitDB()

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
