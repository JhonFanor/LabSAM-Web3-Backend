package app

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	engine := gin.Default()

	return engine
}
