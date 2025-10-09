package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type TypeOfLawService interface {
	GetAllTypeOfLaw() ([]models.TypeOfLaw, error)
}

type typeOfLawService struct {
	repo repositories.TypeOfLawRepository
	db   *gorm.DB
}

func NewTyepOfLawServiceService(repo repositories.TypeOfLawRepository, db *gorm.DB) TypeOfLawService {
	return &typeOfLawService{
		repo: repo,
		db:   db,
	}
}

func (s *typeOfLawService) GetAllTypeOfLaw() ([]models.TypeOfLaw, error) {
	return s.repo.GetAll()
}
