package routes

import (
	"lamsam-web3-backend/internal/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type TypeEducationRoutesParams struct {
	fx.In
	Router                  *gin.Engine
	TypeEducationController *controllers.TypeEducationController
}

type TypeEducationRoutes struct {
	Router                  *gin.Engine
	TypeEducationController *controllers.TypeEducationController
}

func NewTypeEducationRoutes(p TypeEducationRoutesParams) *TypeEducationRoutes {
	return &TypeEducationRoutes{
		Router:                  p.Router,
		TypeEducationController: p.TypeEducationController,
	}
}

func (tr *TypeEducationRoutes) Routes() {
	typeEducation := tr.Router.Group("/api/type-education")
	{
		typeEducation.GET("", tr.TypeEducationController.GetAllTypeEducation)
	}
}
