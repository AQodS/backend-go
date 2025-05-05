package controllers

import (
	"backend-go/database"
	"backend-go/helpers"
	"backend-go/models"
	"backend-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	//init struct untuk menangkap data request
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	// validasi input dari request dengan ShouldBindJSON
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// find user by username, if not found return unauthorized
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// if login success
	token := helpers.GenerateToken(user.Username)

	// send success response
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})
}
