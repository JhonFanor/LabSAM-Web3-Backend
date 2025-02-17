package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type JobExchangeSubtopicRoutesParams struct {
	fx.In
	Router                        *gin.Engine
	ValidatorMiddleware           *middlewares.ValidatorMiddleware
	TonkenMiddleware              *middlewares.TokenMiddleware
	JobExchangeSubtopicController *controllers.JobExchangeSubtopicController
}

type JobExchangeSubtopicRoutes struct {
	Router                        *gin.Engine
	ValidatorMiddleware           *middlewares.ValidatorMiddleware
	TonkenMiddleware              *middlewares.TokenMiddleware
	JobExchangeSubtopicController *controllers.JobExchangeSubtopicController
}

func NewJobExchangeSubtopicRoutes(p JobExchangeSubtopicRoutesParams) *JobExchangeSubtopicRoutes {
	return &JobExchangeSubtopicRoutes{
		Router:                        p.Router,
		ValidatorMiddleware:           p.ValidatorMiddleware,
		TonkenMiddleware:              p.TonkenMiddleware,
		JobExchangeSubtopicController: p.JobExchangeSubtopicController,
	}
}

func (jr *JobExchangeSubtopicRoutes) Routes() {
	jobExchangeSubtopic := jr.Router.Group("/job-exchange-subtopic")
	{
		jobExchangeSubtopic.POST("/create/:id", jr.ValidatorMiddleware.ValidateInput(&requests.JobExchangeSubtopicRequest{}), jr.TonkenMiddleware.ValidateToken(), jr.JobExchangeSubtopicController.CreateJobExchangeSubtopic)
		jobExchangeSubtopic.DELETE("/delete/:id", jr.ValidatorMiddleware.ValidateInput(&requests.JobExchangeSubtopicRequest{}), jr.TonkenMiddleware.ValidateToken(), jr.JobExchangeSubtopicController.DeleteJobExchangeSubtopic)
	}
}
