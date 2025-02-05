package provide

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

func RouteProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			routes.NewAuthRoutes,
		),
	)
}
