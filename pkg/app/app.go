package app

import (
	"github.com/Damillora/Yuriko/pkg/config"
	"github.com/Damillora/Yuriko/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (a *App) Initialize() {
	config.InitializeConfig()
	database.Initialize(config.CurrentConfig)
}

func (a *App) StartApp() {
	g := gin.Default()
	g.Use(cors.Default())
	InitializeRoutes(g)
	g.Run()
}
