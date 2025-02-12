package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type CompanySubtopicService interface {
	CreateCompanySubtopic(companySubtopic *models.CompanySubtopic) (*models.CompanySubtopic, error)
	DeleteCompanySubtopic(companyID, subtopicID uint) error
}

type companySubtopicService struct {
	repo repositories.CompanySubtopicRepository
	db   *gorm.DB
}

func NewCompanySubtopicService(repo repositories.CompanySubtopicRepository, db *gorm.DB) CompanySubtopicService {
	return &companySubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *companySubtopicService) CreateCompanySubtopic(companySubtopic *models.CompanySubtopic) (*models.CompanySubtopic, error) {
	if companySubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(companySubtopic)
}

func (s *companySubtopicService) DeleteCompanySubtopic(companyID, subtopicID uint) error {
	if companyID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(companyID, subtopicID)
}
