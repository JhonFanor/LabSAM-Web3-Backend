package provide

import (
	"lamsam-web3-backend/internal/services"

	"go.uber.org/fx"
)

func ServiceProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			services.NewAuthService,
			services.NewBusinessUserService,
			services.NewPermissionRoleService,
			services.NewPermissionService,
			services.NewPermissionUserService,
			services.NewRegularUserService,
			services.NewRoleService,
			services.NewUniversityUserService,
			services.NewUserService,
		),
	)
}
