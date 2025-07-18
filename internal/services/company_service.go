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
	GetAllCompaniesByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllCompaniesNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountCompaniesNotApproved(role string) (int64, error)
	GetCompanyByID(id uint, userID uint, role string) (*models.Company, error)
	UpdateCompany(company *models.Company, userID uint, role string) error
	SetCompanyApproval(id uint, approved bool, role string) error
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

func (s *companyService) GetAllCompaniesByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *companyService) GetAllCompaniesNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *companyService) CountCompaniesNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *companyService) GetCompanyByID(id uint, userID uint, role string) (*models.Company, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	company, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if company.IsApproved != nil && *company.IsApproved {
		return company, nil
	}

	if company.UserID == userID || role == "admin" {
		return company, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *companyService) UpdateCompany(company *models.Company, userID uint, role string) error {
	if company == nil || company.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetCompanyByID(company.ID, userID, role)
	if err != nil {
		return err
	}

	company.IsApproved = nil

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(company)
	if len(updates) == 0 {
		return nil
	}

	existing.User = nil
	err = s.repo.Update(existing, updates)

	if err == nil && company.LocalitationID != nil && existing.LocalitationID != nil && *company.LocalitationID != *existing.LocalitationID {
		s.localitationService.DeleteLocalitation(*existing.LocalitationID)
	}

	return err
}

func (s *companyService) SetCompanyApproval(id uint, approved bool, role string) error {
	if role != "admin" {
		return customerrors.ErrUnauthorized
	}

	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	isApproved := approved
	updates := map[string]interface{}{
		"is_approved": &isApproved,
	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *companyService) DeleteCompany(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetCompanyByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
