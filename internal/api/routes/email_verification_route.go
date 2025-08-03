package routes

import (
	"lamsam-web3-backend/internal/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type EmailVerificationRoutesParams struct {
	fx.In
	Router                      *gin.Engine
	EmailVerificationController *controllers.EmailVerificationController
}

type EmailVerificationRoutes struct {
	Router                      *gin.Engine
	EmailVerificationController *controllers.EmailVerificationController
}

func NewEmailVerificationRoutes(p EmailVerificationRoutesParams) *EmailVerificationRoutes {
	return &EmailVerificationRoutes{
		Router:                      p.Router,
		EmailVerificationController: p.EmailVerificationController,
	}
}

func (r *EmailVerificationRoutes) Routes() {
	email := r.Router.Group("/api/verify-email")
	{
		email.GET("", r.EmailVerificationController.VerifyEmail)
	}
}
