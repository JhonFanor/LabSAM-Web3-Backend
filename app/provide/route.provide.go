package provide

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

func RouteProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			routes.NewAuthRoutes,
			routes.NewBankOfResumeRoutes,
			routes.NewBankOfResumeSubtopicRoutes,
			routes.NewCompanyRoutes,
			routes.NewCompanySubtopicRoutes,
			routes.NewDocumentationRoutes,
			routes.NewDocumentationSubtopicRoutes,
			routes.NewEducationalOfferRoutes,
			routes.NewEducationalOfferSubtopicRoutes,
			routes.NewEventRoutes,
			routes.NewEventSubtopicRoutes,
			routes.NewInvestigationRoutes,
			routes.NewInvestigationSubtopicRoutes,
			routes.NewJobExchangeRoutes,
			routes.NewJobExchangeSubtopicRoutes,
			routes.NewLegislationRoutes,
			routes.NewLegislationSubtopicRoutes,
			routes.NewNewsRoutes,
			routes.NewNewsSubtopicRoutes,
		),
	)
}
