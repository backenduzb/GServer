package routes

import (
	"github.com/gin-gonic/gin"
	"app/internal/handlers/auth"
	"app/internal/middleware"
)

func SetupAuthRoutes(r *gin.Engine) {
	api := r.Group("/auth")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	api.GET("/me",
		middleware.AuthMiddleware(),
		handlers.Profile,
	)
}