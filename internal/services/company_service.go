package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type CompanyService interface {
	CreateCompany(company *models.Company, userID uint, subtopicIDs []uint) (*models.Company, error)
	GetAllCompanies(c *gin.Context) (*dto.PaginationDTO, error)
	GetCompanyByID(id uint) (*models.Company, error)
	UpdateCompany(company *models.Company, userID uint, role string) error
	DeleteCompany(id uint, userID uint, role string) error
}

type companyService struct {
	repo                   repositories.CompanyRepository
	localitationService    LocalitationService
	companySubtopicService CompanySubtopicService
}

func NewCompanyService(repo repositories.CompanyRepository, localitationService LocalitationService, companySubtopicService CompanySubtopicService) CompanyService {
	return &companyService{
		repo:                   repo,
		localitationService:    localitationService,
		companySubtopicService: companySubtopicService,
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

func (s *companyService) GetAllCompanies(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *companyService) GetCompanyByID(id uint) (*models.Company, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *companyService) UpdateCompany(company *models.Company, userID uint, role string) error {
	if company == nil || company.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetCompanyByID(company.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(company)
	if len(updates) == 0 {
		return nil
	}

	err = s.repo.Update(existing, updates)

	if err == nil && company.LocalitationID != nil && existing.LocalitationID != nil && *company.LocalitationID != *existing.LocalitationID {
		s.localitationService.DeleteLocalitation(*existing.LocalitationID)
	}

	return err
}

func (s *companyService) DeleteCompany(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetCompanyByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
