package routes

import (
	"lamsam-web3-backend/internal/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UniversityTypeRoutesParams struct {
	fx.In
	Router                   *gin.Engine
	UniversityTypeController *controllers.UniversityTypeController
}

type UniversityTypeRoutes struct {
	Router                   *gin.Engine
	UniversityTypeController *controllers.UniversityTypeController
}

func NewUniversityTypeRoutes(p UniversityTypeRoutesParams) *UniversityTypeRoutes {
	return &UniversityTypeRoutes{
		Router:                   p.Router,
		UniversityTypeController: p.UniversityTypeController,
	}
}

func (utr *UniversityTypeRoutes) Routes() {
	group := utr.Router.Group("/api/university-types")
	{
		group.GET("", utr.UniversityTypeController.GetAllUniversityTypes)
	}
}
