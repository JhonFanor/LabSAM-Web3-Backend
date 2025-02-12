package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type EventSubtopicRepository interface {
	Create(eventSubtopic *models.EventSubtopic) (*models.EventSubtopic, error)
	Delete(eventID, subtopicID uint) error
}

type eventSubtopicRepository struct {
	db *gorm.DB
}

func NewEventSubtopicRepository(db *gorm.DB) EventSubtopicRepository {
	return &eventSubtopicRepository{
		db: db,
	}
}

func (r *eventSubtopicRepository) Create(eventSubtopic *models.EventSubtopic) (*models.EventSubtopic, error) {
	if err := r.db.Create(eventSubtopic).Error; err != nil {
		return nil, err
	}
	return eventSubtopic, nil
}

func (r *eventSubtopicRepository) Delete(eventID, subtopicID uint) error {
	return r.db.Where("event_id = ? AND subtopic_id = ?", eventID, subtopicID).
		Delete(&models.EventSubtopic{}).Error
}
