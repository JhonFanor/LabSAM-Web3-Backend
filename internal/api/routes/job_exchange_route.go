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
	jobExchange := jer.Router.Group("/job-exchange")
	{
		jobExchange.POST("/create", jer.ValidatorMiddleware.ValidateInput(&requests.JobExchangeRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobExchangeController.CreateJobExchange)
		jobExchange.GET("/get/all", jer.JobExchangeController.GetAllJobsExchange)
		jobExchange.GET("/get/:id", jer.JobExchangeController.GetJobExchangeByID)
		jobExchange.PUT("/update/:id", jer.ValidatorMiddleware.ValidateInput(&requests.JobExchangeUpdateRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobExchangeController.UpdateJobExchange)
		jobExchange.DELETE("/delete/:id", jer.ValidatorMiddleware.ValidateInput(&requests.JobExchangeUpdateRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobExchangeController.DeleteJobExchange)
	}
}
