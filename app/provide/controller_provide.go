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
			controllers.NewBankOfResumeSubtopicController,
			controllers.NewCompanyController,
			controllers.NewCompanySubtopicController,
			controllers.NewDocumentationController,
			controllers.NewDocumentationSubtopicController,
			controllers.NewEducationalOfferController,
			controllers.NewEducationalOfferSubtopicController,
			controllers.NewEventController,
			controllers.NewEventSubtopicController,
			controllers.NewInvestigationController,
			controllers.NewInvestigationSubtopicController,
			controllers.NewJobExchangeController,
			controllers.NewJobExchangeSubtopicController,
			controllers.NewLegislationController,
			controllers.NewLegislationSubtopicController,
			controllers.NewNewsController,
			controllers.NewNewsSubtopicController,
			controllers.NewUserController,
		),
	)
}
