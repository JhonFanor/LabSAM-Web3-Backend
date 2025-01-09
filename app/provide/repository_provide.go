package provide

import (
	"lamsam-web3-backend/internal/repositories"

	"go.uber.org/fx"
)

func RepositoryProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			repositories.NewUserRepository,
		),
	)
}
