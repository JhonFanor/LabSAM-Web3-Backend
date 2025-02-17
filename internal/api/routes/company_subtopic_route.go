package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type CompanySubtopicRoutesParams struct {
	fx.In
	Router                    *gin.Engine
	ValidatorMiddleware       *middlewares.ValidatorMiddleware
	TonkenMiddleware          *middlewares.TokenMiddleware
	CompanySubtopicController *controllers.CompanySubtopicController
}

type CompanySubtopicRoutes struct {
	Router                    *gin.Engine
	ValidatorMiddleware       *middlewares.ValidatorMiddleware
	TonkenMiddleware          *middlewares.TokenMiddleware
	CompanySubtopicController *controllers.CompanySubtopicController
}

func NewCompanySubtopicRoutes(p CompanySubtopicRoutesParams) *CompanySubtopicRoutes {
	return &CompanySubtopicRoutes{
		Router:                    p.Router,
		ValidatorMiddleware:       p.ValidatorMiddleware,
		TonkenMiddleware:          p.TonkenMiddleware,
		CompanySubtopicController: p.CompanySubtopicController,
	}
}

func (cr *CompanySubtopicRoutes) Routes() {
	companySubtopic := cr.Router.Group("/company-subtopic")
	{
		companySubtopic.POST("/create/:id", cr.ValidatorMiddleware.ValidateInput(&requests.CompanySubtopicRequest{}), cr.TonkenMiddleware.ValidateToken(), cr.CompanySubtopicController.CreateCompanySubtopic)
		companySubtopic.DELETE("/delete/:id", cr.ValidatorMiddleware.ValidateInput(&requests.CompanySubtopicRequest{}), cr.TonkenMiddleware.ValidateToken(), cr.CompanySubtopicController.DeleteCompanySubtopic)
	}
}
