package handlers

import (
	"net/http"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateUser(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(200, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUser(input.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if !services.CheckPassword(user.Password, input.Password) {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
