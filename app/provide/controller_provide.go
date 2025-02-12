package provide

import (
	"lamsam-web3-backend/internal/api/controllers"

	"go.uber.org/fx"
)

func ControllerProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			controllers.NewAuthController,
			controllers.NewBankOfResumeController,
			controllers.NewCompanyController,
			controllers.NewDocumentationController,
			controllers.NewEducationalOfferController,
			controllers.NewEventController,
			controllers.NewInvestigationController,
			controllers.NewJobExchangeController,
			controllers.NewLegislationController,
			controllers.NewNewsController,
		),
	)
}
