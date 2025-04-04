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
			services.NewNewsService,
			services.NewNewsSubtopicService,
			services.NewPermissionRoleService,
			services.NewPermissionService,
			services.NewPermissionUserService,
			services.NewRegularUserService,
			services.NewRoleService,
			services.NewTopicService,
			services.NewUniversityUserService,
			services.NewUserService,
		),
	)
}
