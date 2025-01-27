package provide

import (
	"lamsam-web3-backend/internal/services"

	"go.uber.org/fx"
)

func ServiceProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			services.NewAuthService,
			services.NewUserService,
		),
	)
}
