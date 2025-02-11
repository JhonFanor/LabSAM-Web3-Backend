package app

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	AuthRoutes         *routes.AuthRoutes
	BankOfResumeRoutes *routes.BankOfResumeRoutes
}

func RoutesRegister(p RoutesRegisterParams) {
	p.AuthRoutes.Routes()
	p.BankOfResumeRoutes.Routes()
}
