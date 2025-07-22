package observers

import (
	"log"

	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
)

type AdminNotificationObserver struct {
	notificationRepository repositories.NotificationRepository
	userRepository         repositories.UserRepository
}

func NewAdminNotificationObserver(notificationRepository repositories.NotificationRepository, userRepo repositories.UserRepository) *AdminNotificationObserver {
	return &AdminNotificationObserver{
		notificationRepository: notificationRepository,
		userRepository:         userRepo,
	}
}

func (o *AdminNotificationObserver) Handle(event EventObserver) {
	if event.Type != EventObserverType(Created) && event.Type != EventObserverType(Updated) {
		return
	}

	admins, err := o.userRepository.GetAdmins()
	if err != nil {
		log.Printf("Error obteniendo administradores: %v", err)
		return
	}

	for _, admin := range admins {
		noti := &models.Notification{
			SenderID:     &event.SenderID,
			ReceiverID:   admin.ID,
			Message:      event.Message,
			Action:       event.Action,
			ResourceID:   &event.ResourceID,
			ResourceType: &event.ResourceType,
			IsRead:       false,
		}

		if _, err := o.notificationRepository.Create(noti); err != nil {
			log.Printf("Error creando notificación para admin ID %d: %v", admin.ID, err)
		}
	}
}
