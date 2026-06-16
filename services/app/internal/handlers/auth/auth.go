package handlers

import (
	"app/internal/schemas/user/requests"
	"app/internal/services/auth"
	"app/internal/services/jwt"
	"app/internal/services/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Register(c *gin.Context) {
	var input requests.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := auth.CreateUser(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(200, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var input requests.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.GetUser(input.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if !auth.CheckPassword(user.Password, input.Password) {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token, err := jwt.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, responses.NewJWTTokenResponse(token))
}
