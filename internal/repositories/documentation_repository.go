package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type DocumentationRepository interface {
	Create(documentation *models.Documentation) (*models.Documentation, error)
}

type documentationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewDocumentationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) DocumentationRepository {
	return &documentationRepository{
		db: db,
		qm: qm,
	}
}

func (r *documentationRepository) Create(documentation *models.Documentation) (*models.Documentation, error) {
	if err := r.db.Create(documentation).Error; err != nil {
		return nil, err
	}
	return documentation, nil
}
