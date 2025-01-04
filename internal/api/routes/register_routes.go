package routes

import (
	"LABSAM-WEB3-BACKEND/config"
	"LABSAM-WEB3-BACKEND/internal/logging"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterRoutesParams struct {
	fx.In
	Config *config.AppConfig
	Logger logging.Logger
}

func RegisterRoutes(p RegisterRoutesParams) *http.Server {
	if p.Config.LogLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	p.Logger.LogInfo("Registering basic API routes...")

	router := gin.Default()

	publicRoutes := router.Group("/api")
	{
		// Healthcheck route
		publicRoutes.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		})
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", p.Config.ServerPort),
		Handler: router,
	}

	p.Logger.LogInfo("Basic routes registered successfully")

	return srv
}
