package provide

import (
	"lamsam-web3-backend/internal/services"

	"go.uber.org/fx"
)

func ServiceProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			services.NewAuthService,
			services.NewBankOfResumeService,
			services.NewBankOfResumeSubtopicService,
			services.NewBusinessUserService,
			services.NewCompanyService,
			services.NewCompanySubtopicService,
			services.NewContactService,
			services.NewDeniedPermissionUserService,
			services.NewDocumentationService,
			services.NewDocumentationSubtopicService,
			services.NewEducationalOfferService,
			services.NewEducationalOfferSubtopicService,
			services.NewEventService,
			services.NewEventSubtopicService,
			services.NewInvestigationService,
			services.NewInvestigationSubtopicService,
			services.NewJobBoardService,
			services.NewJobBoardSubtopicService,
			services.NewLegislationService,
			services.NewLegislationSubtopicService,
			services.NewLocalitationService,
			services.NewLocationService,
			services.NewNewsService,
			services.NewNewsSubtopicService,
			services.NewNotificationService,
			services.NewPermissionRoleService,
			services.NewPermissionService,
			services.NewPermissionUserService,
			services.NewPublicationService,
			services.NewRegularUserService,
			services.NewRejectionCommentService,
			services.NewRoleService,
			services.NewSubtopicService,
			services.NewTopicService,
			services.NewUniversityTypeService,
			services.NewUniversityUserService,
			services.NewUserService,
		),
	)
}
