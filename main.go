package main

import (
	"backend-go/config"
	"backend-go/database"
	"backend-go/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	database.InitDB()

	// setup router
	r := routes.SetupRouter()

	// Run the server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
