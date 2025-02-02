package provide

import (
	"lamsam-web3-backend/internal/api/middlewares"

	"go.uber.org/fx"
)

func MiddlewareProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			middlewares.NewRefreshTokenMiddleware,
			middlewares.NewTokenMiddleware,
			middlewares.NewValidatorMiddleware,
		),
	)
}
