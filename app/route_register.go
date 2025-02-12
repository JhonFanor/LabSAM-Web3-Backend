package app

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	AuthRoutes             *routes.AuthRoutes
	BankOfResumeRoutes     *routes.BankOfResumeRoutes
	CompanyRoutes          *routes.CompanyRoutes
	DocumentationRoutes    *routes.DocumentationRoutes
	EducationalOfferRoutes *routes.EducationalOfferRoutes
	EventRoutes            *routes.EventRoutes
	InvestigationRoutes    *routes.InvestigationRoutes
	JobExchangeRoutes      *routes.JobExchangeRoutes
	LegislationRoutes      *routes.LegislationRoutes
	NewsRoutes             *routes.NewsRoutes
}

func RoutesRegister(p RoutesRegisterParams) {
	p.AuthRoutes.Routes()
	p.BankOfResumeRoutes.Routes()
	p.CompanyRoutes.Routes()
	p.DocumentationRoutes.Routes()
	p.EducationalOfferRoutes.Routes()
	p.EventRoutes.Routes()
	p.InvestigationRoutes.Routes()
	p.JobExchangeRoutes.Routes()
	p.LegislationRoutes.Routes()
	p.NewsRoutes.Routes()
}
