package game

import (
	"app/internal/database"
	"app/internal/models"
)

func UpdatePosition(userID uint, x, y, z float64) error {
	result := database.DB.
		Model(&models.Position{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"x": x,
			"y": y,
			"z": z,
		})
	return result.Error
}