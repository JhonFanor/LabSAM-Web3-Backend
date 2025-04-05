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
	TokenMiddleware        *middlewares.TokenMiddleware
	BankOfResumeController *controllers.BankOfResumeController
}

type BankOfResumeRoutes struct {
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	TokenMiddleware        *middlewares.TokenMiddleware
	BankOfResumeController *controllers.BankOfResumeController
}

func NewBankOfResumeRoutes(p BankOfResumeRoutesParams) *BankOfResumeRoutes {
	return &BankOfResumeRoutes{
		Router:                 p.Router,
		ValidatorMiddleware:    p.ValidatorMiddleware,
		TokenMiddleware:        p.TokenMiddleware,
		BankOfResumeController: p.BankOfResumeController,
	}
}

func (br *BankOfResumeRoutes) Routes() {
	bankOfResume := br.Router.Group("/api/bank-of-resume")
	{
		bankOfResume.POST("", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeCreateRequest{}), br.TokenMiddleware.ValidateToken(), br.BankOfResumeController.CreateBankOfResume)
		bankOfResume.GET("", br.BankOfResumeController.GetAllBankOfResumes)
		bankOfResume.GET("/:id", br.BankOfResumeController.GetBankOfResumeByID)
		bankOfResume.PUT("/:id", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeUpdateRequest{}), br.TokenMiddleware.ValidateToken(), br.BankOfResumeController.UpdateBankOfResume)
		bankOfResume.DELETE("/:id", br.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeUpdateRequest{}), br.TokenMiddleware.ValidateToken(), br.BankOfResumeController.DeleteBankOfResume)
	}
}
