package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type LegislationRepository interface {
	Create(legislation *models.Legislation) (*models.Legislation, error)
}

type legislationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewLegislationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) LegislationRepository {
	return &legislationRepository{
		db: db,
		qm: qm,
	}
}

func (r *legislationRepository) Create(legislation *models.Legislation) (*models.Legislation, error) {
	if err := r.db.Create(legislation).Error; err != nil {
		return nil, err
	}
	return legislation, nil
}
