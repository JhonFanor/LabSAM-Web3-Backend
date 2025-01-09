package app

import (
	"lamsam-web3-backend/internal/logging"

	"github.com/gin-gonic/gin"
)

func StartServer(logger logging.Logger, router *gin.Engine) {
	logger.LogInfo("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		logger.LogError("Failed to start server", err)
	}
}
