package routes

import (
	"backend-go/controllers"
	"backend-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	return router
}
