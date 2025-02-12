package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type InvestigationRepository interface {
	Create(investigation *models.Investigation) (*models.Investigation, error)
}

type investigationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewInvestigationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) InvestigationRepository {
	return &investigationRepository{
		db: db,
		qm: qm,
	}
}

func (r *investigationRepository) Create(investigation *models.Investigation) (*models.Investigation, error) {
	if err := r.db.Create(investigation).Error; err != nil {
		return nil, err
	}
	return investigation, nil
}
