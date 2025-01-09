package app

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/api/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	AuthController      *controllers.AuthController
}

func RoutesRegister(p RoutesRegisterParams) {
	routes.AuthRoutes(p.Router, p.ValidatorMiddleware, p.AuthController)
}
