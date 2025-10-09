package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type TypeEducationService interface {
	GetAllTypeEducation() ([]models.TypeEducation, error)
}

type typeEducationService struct {
	repo repositories.TypeEducationRepository
	db   *gorm.DB
}

func NewTypeEducationService(repo repositories.TypeEducationRepository, db *gorm.DB) TypeEducationService {
	return &typeEducationService{
		repo: repo,
		db:   db,
	}
}

func (s *typeEducationService) GetAllTypeEducation() ([]models.TypeEducation, error) {
	return s.repo.GetAll()
}
