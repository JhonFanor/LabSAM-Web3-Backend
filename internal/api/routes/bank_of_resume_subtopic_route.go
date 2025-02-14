package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type BankOfResumeRoutesSubtopicParams struct {
	fx.In
	Router                         *gin.Engine
	ValidatorMiddleware            *middlewares.ValidatorMiddleware
	TonkenMiddleware               *middlewares.TokenMiddleware
	BankOfResumeSubtopicController *controllers.BankOfResumeSubtopicController
}

type BankOfResumeSubtopicRoutes struct {
	Router                         *gin.Engine
	ValidatorMiddleware            *middlewares.ValidatorMiddleware
	TonkenMiddleware               *middlewares.TokenMiddleware
	BankOfResumeSubtopicController *controllers.BankOfResumeSubtopicController
}

func NewBankOfResumeSutopicRoutes(p BankOfResumeRoutesSubtopicParams) *BankOfResumeSubtopicRoutes {
	return &BankOfResumeSubtopicRoutes{
		Router:                         p.Router,
		ValidatorMiddleware:            p.ValidatorMiddleware,
		TonkenMiddleware:               p.TonkenMiddleware,
		BankOfResumeSubtopicController: p.BankOfResumeSubtopicController,
	}
}

func (br *BankOfResumeSubtopicRoutes) Routes() {
	bankOfResumeSubtopic := br.Router.Group("/bank-of-resume-subtopic")
	{
		bankOfResumeSubtopic.POST("/create", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeSubtopicRequest{}), br.TonkenMiddleware.ValidateToken(), br.BankOfResumeSubtopicController.CreateBankOfResumeSubtopic)
		bankOfResumeSubtopic.DELETE("/delete/:id", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeSubtopicRequest{}), br.TonkenMiddleware.ValidateToken(), br.BankOfResumeSubtopicController.DeleteBankOfResumeSubtopic)
	}
}
