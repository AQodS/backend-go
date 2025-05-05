package controllers

import (
	"backend-go/database"
	"backend-go/models"
	"backend-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	// init slice to store user data
	var users []models.User

	// select user data from database
	database.DB.Find(&users)

	// send success response
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Users",
		Data:    users,
	})
}
