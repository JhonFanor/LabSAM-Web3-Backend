package provide

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

func RouteProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			routes.NewAuthRoutes,
			routes.NewBankOfResumeRoutes,
			routes.NewBankOfResumeSubtopicRoutes,
			routes.NewCompanyRoutes,
			routes.NewCompanySubtopicRoutes,
			routes.NewDeniedPermissionUserRoutes,
			routes.NewDocumentationRoutes,
			routes.NewDocumentationSubtopicRoutes,
			routes.NewEducationalOfferRoutes,
			routes.NewEducationalOfferSubtopicRoutes,
			routes.NewEmailVerificationRoutes,
			routes.NewEventRoutes,
			routes.NewEventSubtopicRoutes,
			routes.NewInvestigationRoutes,
			routes.NewInvestigationSubtopicRoutes,
			routes.NewJobBoardRoutes,
			routes.NewJobBoardSubtopicRoutes,
			routes.NewLegislationRoutes,
			routes.NewLegislationSubtopicRoutes,
			routes.NewNewsRoutes,
			routes.NewNewsSubtopicRoutes,
			routes.NewNotificationRoutes,
			routes.NewPermissionRoleRoutes,
			routes.NewPermissionRoutes,
			routes.NewPermissionUserRoutes,
			routes.NewPublicationRoutes,
			routes.NewRejectionCommentRoutes,
			routes.NewTopicRoutes,
			routes.NewUniversityTypeRoutes,
			routes.NewUploadRoutes,
			routes.NewUserRoutes,
		),
	)
}
