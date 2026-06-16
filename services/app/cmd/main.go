package main

import (
	"github.com/gin-gonic/gin"
	"app/internal/routes"
	"app/internal/database"
	"app/config/settings"
)

func main() {
	settings.LoadEnv()
	router := gin.Default()
	
	database.Connect(settings.Envs.DB_URL)
	routes.SetupAuthRoutes(router)
	routes.SetupPositionRoutes(router)
	router.Run(settings.Envs.PORT)
}