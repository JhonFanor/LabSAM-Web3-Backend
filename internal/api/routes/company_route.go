package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type CompanyRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	CompanyController   *controllers.CompanyController
}

type CompanyRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	CompanyController   *controllers.CompanyController
}

func NewCompanyRoutes(p CompanyRoutesParams) *CompanyRoutes {
	return &CompanyRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		CompanyController:   p.CompanyController,
	}
}

func (cr *CompanyRoutes) Routes() {
	company := cr.Router.Group("/company")
	{
		company.POST("/create", cr.ValidatorMiddleware.ValidateInput(&requests.CompanyRequest{}), cr.TokenMiddleware.ValidateToken(), cr.CompanyController.CreateCompany)
		company.GET("/get/all", cr.CompanyController.GetAllCompanies)
		company.GET("/get/:id", cr.CompanyController.GetCompanyByID)
		company.PUT("/update/:id", cr.ValidatorMiddleware.ValidateInput(&requests.CompanyUpdateRequest{}), cr.TokenMiddleware.ValidateToken(), cr.CompanyController.UpdateCompany)
		company.DELETE("/delete/:id", cr.ValidatorMiddleware.ValidateInput(&requests.CompanyUpdateRequest{}), cr.TokenMiddleware.ValidateToken(), cr.CompanyController.DeleteCompany)
	}
}
