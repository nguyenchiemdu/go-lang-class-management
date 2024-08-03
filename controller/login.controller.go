package controller

import (
	"http_request/class-management/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) HandleLogin(c *gin.Context) {

	var request LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username, password are required"})
		return
	}

	username := request.Username
	password := request.Password

	// Get user from Database
	user, err := ctrl.DatabaseService.GetUserByUsername(username)

	if err != nil {
		//TODO: should create a general format for responses
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	} else if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !utils.ValidateHashPasword(password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.CreateAccessToken(map[string]interface{}{
		"key": username}, time.Duration(ctrl.appConfig.JWTExpire), []byte(ctrl.appConfig.JWTSecret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
