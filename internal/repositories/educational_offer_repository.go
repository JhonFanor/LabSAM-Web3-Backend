package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type EducationalOfferRepository interface {
	Create(educationalOffer *models.EducationalOffer) (*models.EducationalOffer, error)
}

type educationalOfferRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewEducationalOfferRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) EducationalOfferRepository {
	return &educationalOfferRepository{
		db: db,
		qm: qm,
	}
}

func (r *educationalOfferRepository) Create(educationalOffer *models.EducationalOffer) (*models.EducationalOffer, error) {
	if err := r.db.Create(educationalOffer).Error; err != nil {
		return nil, err
	}
	return educationalOffer, nil
}
