package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

type NotificationService interface {
	GetAllNotifications(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	CountNotificationNotRead(userID uint) (int64, error)
	GetNotificationByID(id uint, userID uint) (*models.Notification, error)
	UpdateIsRead(id uint, userID uint) error
	UpdateAllIsReadByUserID(userID uint) error
	DeleteNotification(id uint, userID uint) error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{
		repo: repo,
	}
}

func (s *notificationService) GetAllNotifications(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	if userID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetAll(c, userID)
}

func (s *notificationService) CountNotificationNotRead(userID uint) (int64, error) {
	if userID == 0 {
		return 0, customerrors.ErrInvalidID
	}
	return s.repo.CountNotRead(userID)
}

func (s *notificationService) GetNotificationByID(id uint, userID uint) (*models.Notification, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	notification, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if notification.ReceiverID != userID {
		return nil, customerrors.ErrUnauthorized
	}

	return notification, nil
}

func (s *notificationService) UpdateIsRead(id uint, userID uint) error {
	notification, err := s.GetNotificationByID(id, userID)
	if err != nil {
		return err
	}

	if notification.IsRead {
		return nil
	}

	return s.repo.UpdateIsRead(id, true)
}

func (s *notificationService) UpdateAllIsReadByUserID(userID uint) error {
	if userID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.UpdateAllIsReadByReceiverID(userID, true)
}

func (s *notificationService) DeleteNotification(id uint, userID uint) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}
	existing, err := s.GetNotificationByID(id, userID)
	if err != nil {
		return err
	}

	if existing.ReceiverID != userID {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
