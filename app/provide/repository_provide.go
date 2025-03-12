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
			repositories.NewDocumentationRepository,
			repositories.NewDocumentationSubtopicRepository,
			repositories.NewEducationalOfferRepository,
			repositories.NewEducationalOfferSubtopicRepository,
			repositories.NewEventRepository,
			repositories.NewEventSubtopicRepository,
			repositories.NewInvestigationRepository,
			repositories.NewInvestigationSubtopicRepository,
			repositories.NewJobExchangeRepository,
			repositories.NewJobExchangeSubtopicRepository,
			repositories.NewLegislationRepository,
			repositories.NewLegislationSubtopicRepository,
			repositories.NewNewsRepository,
			repositories.NewNewsSubtopicRepository,
			repositories.NewPermissionRepository,
			repositories.NewPermissionRoleRepository,
			repositories.NewPermissionUserRepository,
			repositories.NewRegularUserRepository,
			repositories.NewRoleRepository,
			repositories.NewTopicRepository,
			repositories.NewUniversityUserRepository,
			repositories.NewUserRepository,
		),
	)
}
