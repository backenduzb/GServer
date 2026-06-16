package responses

import (
	"app/internal/schemas/user/response"
	"app/internal/models"
)

func NewProfileResponse(user models.User) response.ProfileResponse {
	return response.ProfileResponse{
		ID: int(user.ID),
		Username: user.Username,
		CreatedAt: user.CreatedAt,

		Position: response.PositionResponse{
			X: user.Position.X,
			Y: user.Position.Y,
			Z: user.Position.Z,
		},
	}
} 
