package app

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	AuthRoutes *routes.AuthRoutes
}

func RoutesRegister(p RoutesRegisterParams) {
	p.AuthRoutes.Routes()
}
