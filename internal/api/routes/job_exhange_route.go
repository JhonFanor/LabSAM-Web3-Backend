package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type JobExchangeRoutesParams struct {
	fx.In
	Router                *gin.Engine
	ValidatorMiddleware   *middlewares.ValidatorMiddleware
	TokenMiddleware       *middlewares.TokenMiddleware
	JobExchangeController *controllers.JobExchangeController
}

type JobExchangeRoutes struct {
	Router                *gin.Engine
	ValidatorMiddleware   *middlewares.ValidatorMiddleware
	TokenMiddleware       *middlewares.TokenMiddleware
	JobExchangeController *controllers.JobExchangeController
}

func NewJobExchangeRoutes(p JobExchangeRoutesParams) *JobExchangeRoutes {
	return &JobExchangeRoutes{
		Router:                p.Router,
		ValidatorMiddleware:   p.ValidatorMiddleware,
		TokenMiddleware:       p.TokenMiddleware,
		JobExchangeController: p.JobExchangeController,
	}
}

func (jer *JobExchangeRoutes) Routes() {
	auth := jer.Router.Group("/job-exchange")
	{
		auth.POST("/create", jer.ValidatorMiddleware.ValidateInput(&requests.JobExchangeRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobExchangeController.CreateJobExchange)
	}
}
