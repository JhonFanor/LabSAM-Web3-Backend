package main

import (
	"lamsam-web3-backend/app"
	"lamsam-web3-backend/app/provide"
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/docs"
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/logging"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logging.NewLogger,
			app.NewGinEngine,
			provide.DatabaseProvide,
			config.NewDatabaseConfig,
			gormmanagers.NewDBManager,
		),
		provide.ValidatorProvide(),
		provide.RepositoryProvide(),
		provide.ServiceProvide(),
		provide.MiddlewareProvide(),
		provide.ControllerProvide(),
		fx.Invoke(app.RoutesRegister),
		fx.Invoke(docs.SetupSwagger),
		fx.Invoke(app.StartServer),
	).Run()
}
