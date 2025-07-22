package provide

import (
	"lamsam-web3-backend/internal/observers"
	"lamsam-web3-backend/internal/repositories"

	"go.uber.org/fx"
)

var ObserverProvide = fx.Options(
	fx.Provide(NewAdminNotificationObserver),
)

func NewAdminNotificationObserver(repository repositories.NotificationRepository, userRepo repositories.UserRepository,
) *observers.AdminNotificationObserver {
	return observers.NewAdminNotificationObserver(repository, userRepo)
}
