package provide

import (
	"lamsam-web3-backend/internal/api/controllers"

	"go.uber.org/fx"
)

func ControllerProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			controllers.NewAuthController,
			controllers.NewBankOfResumeController,
			controllers.NewBankOfResumeSubtopicController,
			controllers.NewCompanyController,
			controllers.NewCompanySubtopicController,
			controllers.NewDeniedPermissionUserController,
			controllers.NewDocumentationController,
			controllers.NewDocumentationSubtopicController,
			controllers.NewEducationalOfferController,
			controllers.NewEducationalOfferSubtopicController,
			controllers.NewEmailVerificationController,
			controllers.NewEventController,
			controllers.NewEventSubtopicController,
			controllers.NewInvestigationController,
			controllers.NewInvestigationSubtopicController,
			controllers.NewJobBoardController,
			controllers.NewJobBoardSubtopicController,
			controllers.NewLegislationController,
			controllers.NewLegislationSubtopicController,
			controllers.NewNewsController,
			controllers.NewNewsSubtopicController,
			controllers.NewNotificationController,
			controllers.NewPermissionController,
			controllers.NewPermissionRoleController,
			controllers.NewPermissionUserController,
			controllers.NewPublicationsController,
			controllers.NewRejectionCommentController,
			controllers.NewTopicController,
			controllers.NewTypeEducationController,
			controllers.NewTypeOfLawController,
			controllers.NewUniversityTypeController,
			controllers.NewUploadController,
			controllers.NewUserController,
		),
	)
}
