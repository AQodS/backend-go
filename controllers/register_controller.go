package controllers

import (
	"backend-go/database"
	"backend-go/helpers"
	"backend-go/models"
	"backend-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// init struct untuk menangkap data request
	var req = structs.UserCreateRequest{}

	// validasi request JSON dengan binding Gin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// create new user data with hash password
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	// save new user data to database
	if err := database.DB.Create(&user).Error; err != nil {
		// check if error is duplicate entry error
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			// other error
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	// if success, return success response
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
