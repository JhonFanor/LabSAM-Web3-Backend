package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type InvestigationSubtopicRepository interface {
	Create(investigationSubtopic *models.InvestigationSubtopic) (*models.InvestigationSubtopic, error)
	Delete(investigationID, subtopicID uint) error
}

type investigationSubtopicRepository struct {
	db *gorm.DB
}

func NewInvestigationSubtopicRepository(db *gorm.DB) InvestigationSubtopicRepository {
	return &investigationSubtopicRepository{
		db: db,
	}
}

func (r *investigationSubtopicRepository) Create(investigationSubtopic *models.InvestigationSubtopic) (*models.InvestigationSubtopic, error) {
	if err := r.db.Create(investigationSubtopic).Error; err != nil {
		return nil, err
	}
	return investigationSubtopic, nil
}

func (r *investigationSubtopicRepository) Delete(investigationID, subtopicID uint) error {
	return r.db.Where("investigation_id = ? AND subtopic_id = ?", investigationID, subtopicID).
		Delete(&models.InvestigationSubtopic{}).Error
}
