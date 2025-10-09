package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type TypeOfLawRepository interface {
	GetAll() ([]models.TypeOfLaw, error)
}

type typeOfLawRepository struct {
	db *gorm.DB
}

func NewTypeOfLawRepository(db *gorm.DB) TypeOfLawRepository {
	return &typeOfLawRepository{db: db}
}

func (r *typeOfLawRepository) GetAll() ([]models.TypeOfLaw, error) {
	var typeOfLaw []models.TypeOfLaw

	if err := r.db.Find(&typeOfLaw).Error; err != nil {
		return nil, err
	}

	return typeOfLaw, nil
}
