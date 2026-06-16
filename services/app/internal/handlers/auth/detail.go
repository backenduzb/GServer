package handlers

import (
	"app/internal/services"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"error": "Invalid",
		})
		return
	}

	user, err := services.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User was INVALID",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"id": user.ID,
		"username": user.Username,
		"created_at": user.CreatedAt,
	})
}
