package handlers

import (
	"app/internal/services/auth"
	"github.com/gin-gonic/gin"
	"app/internal/services/responses"
)

func Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"error": "Invalid",
		})
		return
	}

	user, err := auth.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User was INVALID",
		})
		return
	}
	
	c.JSON(200, responses.NewProfileResponse(user))
}
