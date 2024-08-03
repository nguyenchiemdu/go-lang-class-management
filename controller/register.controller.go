package controller

import (
	"fmt"
	"http_request/class-management/models"
	"http_request/class-management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) HandleRegister(c *gin.Context) {
	var request RegisterRequest

	// Bind JSON request to the RegisterRequest struct
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username, password, and email are required"})
		return
	}

	username := request.Username
	password := request.Password

	// Check if user already exists
	existingUser, err := ctrl.DatabaseService.GetUserByUsername(username)
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	} else if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user object
	user := &models.User{
		Key:      username,
		Username: username,
		Password: hashedPassword,
	}

	// Save user to the database
	if err := ctrl.DatabaseService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
