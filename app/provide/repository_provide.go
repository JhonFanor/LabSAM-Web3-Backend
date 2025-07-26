package provide

import (
	"lamsam-web3-backend/internal/repositories"

	"go.uber.org/fx"
)

func RepositoryProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			repositories.NewBankOfResumeRepository,
			repositories.NewBankOfResumeSubtopicRepository,
			repositories.NewBusinessUserRepository,
			repositories.NewCompanyRepository,
			repositories.NewCompanySubtopicRepository,
			repositories.NewContactRepository,
			repositories.NewDeniedPermissionUserRepository,
			repositories.NewDocumentationRepository,
			repositories.NewDocumentationSubtopicRepository,
			repositories.NewEducationalOfferRepository,
			repositories.NewEducationalOfferSubtopicRepository,
			repositories.NewEventRepository,
			repositories.NewEventSubtopicRepository,
			repositories.NewInvestigationRepository,
			repositories.NewInvestigationSubtopicRepository,
			repositories.NewJobBoardRepository,
			repositories.NewJobBoardSubtopicRepository,
			repositories.NewLegislationRepository,
			repositories.NewLegislationSubtopicRepository,
			repositories.NewLocalitationRepository,
			repositories.NewLocationRepository,
			repositories.NewNewsRepository,
			repositories.NewNewsSubtopicRepository,
			repositories.NewNotificationRepository,
			repositories.NewPermissionRepository,
			repositories.NewPermissionRoleRepository,
			repositories.NewPermissionUserRepository,
			repositories.NewRegularUserRepository,
			repositories.NewRoleRepository,
			repositories.NewSubtopicRepository,
			repositories.NewTopicRepository,
			repositories.NewUniversityTypeRepository,
			repositories.NewUniversityUserRepository,
			repositories.NewUserRepository,
		),
	)
}
