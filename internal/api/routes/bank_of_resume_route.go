package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type BankOfResumeRoutesParams struct {
	fx.In
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	TonkenMiddleware       *middlewares.TokenMiddleware
	BankOfResumeController *controllers.BankOfResumeController
}

type BankOfResumeRoutes struct {
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	TonkenMiddleware       *middlewares.TokenMiddleware
	BankOfResumeController *controllers.BankOfResumeController
}

func NewBankOfResumeRoutes(p BankOfResumeRoutesParams) *BankOfResumeRoutes {
	return &BankOfResumeRoutes{
		Router:                 p.Router,
		ValidatorMiddleware:    p.ValidatorMiddleware,
		TonkenMiddleware:       p.TonkenMiddleware,
		BankOfResumeController: p.BankOfResumeController,
	}
}

func (br *BankOfResumeRoutes) Routes() {
	auth := br.Router.Group("/bank-of-resume")
	{
		auth.POST("/create", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeRequest{}), br.TonkenMiddleware.ValidateToken(), br.BankOfResumeController.CreateBankOfResume)
	}
}
