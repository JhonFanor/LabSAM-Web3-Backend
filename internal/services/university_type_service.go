package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
)

type UniversityTypeService interface {
	GetAllUniversityTypes() ([]models.UniversityType, error)
}

type universityTypeService struct {
	repo repositories.UniversityTypeRepository
}

func NewUniversityTypeService(repo repositories.UniversityTypeRepository) UniversityTypeService {
	return &universityTypeService{
		repo: repo,
	}
}

func (s *universityTypeService) GetAllUniversityTypes() ([]models.UniversityType, error) {
	return s.repo.GetAll()
}
