package routes

import (
	"github.com/gin-gonic/gin"
	"app/internal/handlers/position"
	"app/internal/middleware/auth"
)

func SetupPositionRoutes(r *gin.Engine) {
	api := r.Group("/pos")

	api.PUT("/update",
		auth.AuthMiddleware(),
		handlers.UpdatePlayerPosition,
	)
}