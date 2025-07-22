package observers

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"log"
)

type UserNotificationObserver struct {
	notificationRepository repositories.NotificationRepository
}

func NewUserNotificationObserver(notificationRepository repositories.NotificationRepository) *UserNotificationObserver {
	return &UserNotificationObserver{notificationRepository: notificationRepository}
}

func (o *UserNotificationObserver) Handle(event EventObserver) {
	if event.Type != EventObserverType(Approved) && event.Type != EventObserverType(Rejected) {
		return
	}

	noti := &models.Notification{
		SenderID:     &event.SenderID,
		ReceiverID:   event.ReceiverID,
		Message:      event.Message,
		Action:       event.Action,
		ResourceID:   &event.ResourceID,
		ResourceType: &event.ResourceType,
		IsRead:       false,
	}

	_, err := o.notificationRepository.Create(noti)
	if err != nil {
		log.Printf("Error creando notificación para usuario: %v", err)
	}
}
