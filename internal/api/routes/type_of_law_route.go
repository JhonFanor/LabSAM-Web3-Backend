package routes

import (
	"lamsam-web3-backend/internal/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type TypeOfLawRoutesParams struct {
	fx.In
	Router              *gin.Engine
	TypeOfLawController *controllers.TypeOfLawController
}

type TypeOfLawRoutes struct {
	Router              *gin.Engine
	TypeOfLawController *controllers.TypeOfLawController
}

func NewTypeOfLawRoutes(p TypeOfLawRoutesParams) *TypeOfLawRoutes {
	return &TypeOfLawRoutes{
		Router:              p.Router,
		TypeOfLawController: p.TypeOfLawController,
	}
}

func (tr *TypeOfLawRoutes) Routes() {
	topic := tr.Router.Group("/api/type-of-law")
	{
		topic.GET("", tr.TypeOfLawController.GetAllTypeOfLaw)
	}
}
