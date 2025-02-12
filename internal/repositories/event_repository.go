package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) (*models.Event, error)
}

type eventRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewEventRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) EventRepository {
	return &eventRepository{
		db: db,
		qm: qm,
	}
}

func (r *eventRepository) Create(event *models.Event) (*models.Event, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}
