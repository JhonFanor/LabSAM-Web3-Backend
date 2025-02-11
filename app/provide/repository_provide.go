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
			repositories.NewPermissionRepository,
			repositories.NewPermissionRoleRepository,
			repositories.NewPermissionUserRepository,
			repositories.NewRegularUserRepository,
			repositories.NewRoleRepository,
			repositories.NewUniversityUserRepository,
			repositories.NewUserRepository,
		),
	)
}
