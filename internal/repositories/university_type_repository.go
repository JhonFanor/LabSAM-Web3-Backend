package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type UniversityTypeRepository interface {
	GetAll() ([]models.UniversityType, error)
}

type universityTypeRepository struct {
	db *gorm.DB
}

func NewUniversityTypeRepository(db *gorm.DB) UniversityTypeRepository {
	return &universityTypeRepository{
		db: db,
	}
}

func (r *universityTypeRepository) GetAll() ([]models.UniversityType, error) {
	var types []models.UniversityType
	if err := r.db.Find(&types).Error; err != nil {
		return nil, err
	}
	return types, nil
}
