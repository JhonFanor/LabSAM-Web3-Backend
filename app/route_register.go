package app

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	AuthRoutes                  *routes.AuthRoutes
	BankOfResumeRoutes          *routes.BankOfResumeRoutes
	BankOfResumeSubtopicRoutes  *routes.BankOfResumeSubtopicRoutes
	CompanyRoutes               *routes.CompanyRoutes
	CompanySubtopicRoutes       *routes.CompanySubtopicRoutes
	DocumentationRoutes         *routes.DocumentationRoutes
	DocumentationSubtopicRoutes *routes.DocumentationSubtopicRoutes
	EducationalOfferRoutes      *routes.EducationalOfferRoutes
	EducationalOfferSubtopic    *routes.EducationalOfferSubtopicRoutes
	EventRoutes                 *routes.EventRoutes
	EventSubtopicRoutes         *routes.EventSubtopicRoutes
	InvestigationRoutes         *routes.InvestigationRoutes
	InvestigationSubtopicRoutes *routes.InvestigationSubtopicRoutes
	JobBoardRoutes              *routes.JobBoardRoutes
	JobBoardSubtopicRoutes      *routes.JobBoardSubtopicRoutes
	LegislationRoutes           *routes.LegislationRoutes
	LegislationSubtopicRoutes   *routes.LegislationSubtopicRoutes
	NewsRoutes                  *routes.NewsRoutes
	NewsSubtopicRoutes          *routes.NewsSubtopicRoutes
	TopicRoutes                 *routes.TopicRoutes
	UserRoutes                  *routes.UserRoutes
}

func RoutesRegister(p RoutesRegisterParams) {
	p.AuthRoutes.Routes()
	p.BankOfResumeRoutes.Routes()
	p.BankOfResumeSubtopicRoutes.Routes()
	p.CompanyRoutes.Routes()
	p.CompanySubtopicRoutes.Routes()
	p.DocumentationRoutes.Routes()
	p.DocumentationSubtopicRoutes.Routes()
	p.EducationalOfferRoutes.Routes()
	p.EducationalOfferSubtopic.Routes()
	p.EventRoutes.Routes()
	p.EventSubtopicRoutes.Routes()
	p.JobBoardRoutes.Routes()
	p.JobBoardSubtopicRoutes.Routes()
	p.LegislationRoutes.Routes()
	p.LegislationSubtopicRoutes.Routes()
	p.NewsRoutes.Routes()
	p.NewsSubtopicRoutes.Routes()
	p.TopicRoutes.Routes()
	p.UserRoutes.Routes()
}
