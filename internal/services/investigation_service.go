package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InvestigationService interface {
	CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error)
	GetAllInvestigations(c *gin.Context) (*dto.PaginationDTO, error)
	GetInvestigationByID(id uint) (*models.Investigation, error)
	UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error
	DeleteInvestigation(id uint, userID uint, role string) error
}

type investigationService struct {
	repo                         repositories.InvestigationRepository
	investigationSubtopicService InvestigationSubtopicService
	db                           *gorm.DB
}

func NewInvestigationService(repo repositories.InvestigationRepository, investigationSubtopicService InvestigationSubtopicService, db *gorm.DB) InvestigationService {
	return &investigationService{
		repo:                         repo,
		investigationSubtopicService: investigationSubtopicService,
		db:                           db,
	}
}

func (s *investigationService) CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error) {
	if investigation == nil {
		return nil, customerrors.ErrInvalidData
	}

	investigation.UserID = userID

	createdInvestigation, err := s.repo.Create(investigation)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		investigationSubtopic := &models.InvestigationSubtopic{
			InvestigationID: createdInvestigation.ID,
			SubtopicID:      uint(subtopicID),
		}

		_, err := s.investigationSubtopicService.CreateInvestigationSubtopic(investigationSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdInvestigation, nil
}

func (s *investigationService) GetAllInvestigations(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *investigationService) GetInvestigationByID(id uint) (*models.Investigation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *investigationService) UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error {
	if investigation == nil || investigation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetInvestigationByID(investigation.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.GetModifiedFields(existing, investigation)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(investigation)
}

func (s *investigationService) DeleteInvestigation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetInvestigationByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
