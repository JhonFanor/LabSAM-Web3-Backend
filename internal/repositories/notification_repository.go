package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *models.Notification) (*models.Notification, error)
	GetAll(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	CountNotRead(userID uint) (int64, error)
	GetByID(id uint) (*models.Notification, error)
	UpdateIsRead(id uint, isRead bool) error
	UpdateAllIsReadByReceiverID(receiverID uint, isRead bool) error
	Delete(id uint) error
}

type notificationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewNotificationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *notificationRepository) Create(notification *models.Notification) (*models.Notification, error) {
	if err := r.dbManager.Create(notification, r.db); err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *notificationRepository) GetAll(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	query := r.db.
		Preload("Sender").
		Preload("Sender.RegularUser").
		Preload("Receiver").
		Where("receiver_id = ?", userID).
		Order("created_at DESC")

	return r.qm.ApplyPaginationAndFilters(c, query, &models.Notification{}), nil
}

func (r *notificationRepository) CountNotRead(userID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Notification{}).
		Where("receiver_id = ? AND is_read = false", userID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *notificationRepository) GetByID(id uint) (*models.Notification, error) {
	var notification models.Notification
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&notification, conditions, r.db.Preload("Sender").Preload("Receiver")); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) UpdateIsRead(id uint, isRead bool) error {
	updates := map[string]interface{}{"is_read": isRead}
	notification := &models.Notification{ID: id}
	return r.dbManager.Update(notification, updates, r.db)
}

func (r *notificationRepository) UpdateAllIsReadByReceiverID(receiverID uint, isRead bool) error {
	return r.db.
		Model(&models.Notification{}).
		Where("receiver_id = ?", receiverID).
		Update("is_read", isRead).
		Error
}

func (r *notificationRepository) Delete(id uint) error {
	notification := &models.Notification{ID: id}
	return r.dbManager.Delete(notification, r.db)
}
