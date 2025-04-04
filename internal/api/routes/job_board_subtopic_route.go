package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type JobBoardSubtopicRoutesParams struct {
	fx.In
	Router                     *gin.Engine
	ValidatorMiddleware        *middlewares.ValidatorMiddleware
	TonkenMiddleware           *middlewares.TokenMiddleware
	JobBoardSubtopicController *controllers.JobBoardSubtopicController
}

type JobBoardSubtopicRoutes struct {
	Router                     *gin.Engine
	ValidatorMiddleware        *middlewares.ValidatorMiddleware
	TonkenMiddleware           *middlewares.TokenMiddleware
	JobBoardSubtopicController *controllers.JobBoardSubtopicController
}

func NewJobBoardSubtopicRoutes(p JobBoardSubtopicRoutesParams) *JobBoardSubtopicRoutes {
	return &JobBoardSubtopicRoutes{
		Router:                     p.Router,
		ValidatorMiddleware:        p.ValidatorMiddleware,
		TonkenMiddleware:           p.TonkenMiddleware,
		JobBoardSubtopicController: p.JobBoardSubtopicController,
	}
}

func (jr *JobBoardSubtopicRoutes) Routes() {
	jobBoardSubtopic := jr.Router.Group("/job-exchange-subtopic")
	{
		jobBoardSubtopic.POST("/create/:id", jr.ValidatorMiddleware.ValidateInput(&requests.JobBoardSubtopicRequest{}), jr.TonkenMiddleware.ValidateToken(), jr.JobBoardSubtopicController.CreateJobBoardSubtopic)
		jobBoardSubtopic.DELETE("/delete/:id", jr.ValidatorMiddleware.ValidateInput(&requests.JobBoardSubtopicRequest{}), jr.TonkenMiddleware.ValidateToken(), jr.JobBoardSubtopicController.DeleteJobBoardSubtopic)
	}
}
