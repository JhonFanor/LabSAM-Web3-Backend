package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type LegislationSubtopicRepository interface {
	Create(legislationSubtopic *models.LegislationSubtopic) (*models.LegislationSubtopic, error)
	Delete(legislationID, subtopicID uint) error
}

type legislationSubtopicRepository struct {
	db *gorm.DB
}

func NewLegislationSubtopicRepository(db *gorm.DB) LegislationSubtopicRepository {
	return &legislationSubtopicRepository{
		db: db,
	}
}

func (r *legislationSubtopicRepository) Create(legislationSubtopic *models.LegislationSubtopic) (*models.LegislationSubtopic, error) {
	if err := r.db.Create(legislationSubtopic).Error; err != nil {
		return nil, err
	}
	return legislationSubtopic, nil
}

func (r *legislationSubtopicRepository) Delete(legislationID, subtopicID uint) error {
	return r.db.Where("legislation_id = ? AND subtopic_id = ?", legislationID, subtopicID).
		Delete(&models.LegislationSubtopic{}).Error
}
