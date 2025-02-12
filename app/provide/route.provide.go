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
			routes.NewCompanyRoutes,
			routes.NewDocumentationRoutes,
			routes.NewEducationalOfferRoutes,
			routes.NewEventRoutes,
			routes.NewInvestigationRoutes,
			routes.NewJobExchangeRoutes,
			routes.NewLegislationRoutes,
			routes.NewNewsRoutes,
		),
	)
}
