package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type CompanyService interface {
	CreateCompany(company *models.Company, userID uint, subtopicIDs []uint) (*models.Company, error)
}

type companyService struct {
	repo                   repositories.CompanyRepository
	companySubtopicService CompanySubtopicService
	db                     *gorm.DB
	qm                     *gormmanagers.GormQueryManager
}

func NewCompanyService(repo repositories.CompanyRepository, companySubtopicService CompanySubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) CompanyService {
	return &companyService{
		repo:                   repo,
		companySubtopicService: companySubtopicService,
		db:                     db,
		qm:                     qm,
	}
}

func (s *companyService) CreateCompany(company *models.Company, userID uint, subtopicIDs []uint) (*models.Company, error) {
	if company == nil {
		return nil, customerrors.ErrInvalidData
	}

	company.UserID = userID

	createdCompany, err := s.repo.Create(company)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		companySubtopic := &models.CompanySubtopic{
			CompanyID:  createdCompany.ID,
			SubtopicID: uint(subtopicID),
		}

		_, err := s.companySubtopicService.CreateCompanySubtopic(companySubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdCompany, nil
}
