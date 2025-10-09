package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type TypeEducationRepository interface {
	GetAll() ([]models.TypeEducation, error)
}

type typeEducationRepository struct {
	db *gorm.DB
}

func NewTypeEducationRepository(db *gorm.DB) TypeEducationRepository {
	return &typeEducationRepository{db: db}
}

func (r *typeEducationRepository) GetAll() ([]models.TypeEducation, error) {
	var typeEducation []models.TypeEducation

	if err := r.db.Find(&typeEducation).Error; err != nil {
		return nil, err
	}

	return typeEducation, nil
}
