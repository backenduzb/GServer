package handlers

import (
	"app/internal/schemas/position/requests"
	"app/internal/services/game"

	"github.com/gin-gonic/gin"
)

func UpdatePlayerPosition(c* gin.Context) {
	var input requests.UpdatePositionResponse

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	err := game.UpdatePosition(
		userID,
		input.X,
		input.Y,
		input.Z,
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Invalid update function",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "position updated",
	})
}

