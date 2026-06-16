package response

import "time"

type PositionResponse struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type ProfileResponse struct {
	ID int `json:"id"`
	Username string `json:"username"`
	CreatedAt time.Time `json:"created_at"`

	Position PositionResponse `json:"position"` 
}