package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RejectionCommentRoutesParams struct {
	fx.In
	Router                     *gin.Engine
	ValidatorMiddleware        *middlewares.ValidatorMiddleware
	TokenMiddleware            *middlewares.TokenMiddleware
	RejectionCommentController *controllers.RejectionCommentController
}

type RejectionCommentRoutes struct {
	router                     *gin.Engine
	validatorMiddleware        *middlewares.ValidatorMiddleware
	tokenMiddleware            *middlewares.TokenMiddleware
	rejectionCommentController *controllers.RejectionCommentController
}

func NewRejectionCommentRoutes(p RejectionCommentRoutesParams) *RejectionCommentRoutes {
	return &RejectionCommentRoutes{
		router:                     p.Router,
		validatorMiddleware:        p.ValidatorMiddleware,
		tokenMiddleware:            p.TokenMiddleware,
		rejectionCommentController: p.RejectionCommentController,
	}
}

func (r *RejectionCommentRoutes) Routes() {
	rc := r.router.Group("/api/rejection-comments")
	{
		rc.POST("",
			r.validatorMiddleware.ValidateInput(&requests.RejectionCommentCreateRequest{}),
			r.tokenMiddleware.ValidateToken(),
			r.rejectionCommentController.Create,
		)

		rc.PUT("/:id",
			r.validatorMiddleware.ValidateInput(&requests.RejectionCommentUpdateRequest{}),
			r.tokenMiddleware.ValidateToken(),
			r.rejectionCommentController.Update,
		)

		rc.DELETE("/:id",
			r.tokenMiddleware.ValidateToken(),
			r.rejectionCommentController.Delete,
		)

		rc.GET("/:resource_type/:resource_id",
			r.tokenMiddleware.ValidateToken(),
			r.rejectionCommentController.GetAllByResource,
		)
	}
}
