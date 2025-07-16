package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type InvestigationService interface {
	CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error)
	GetAllInvestigations(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllInvestigationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllInvestigationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountInvestigationsNotApproved(role string) (int64, error)
	GetInvestigationByID(id uint, userID uint, role string) (*models.Investigation, error)
	UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error
	DeleteInvestigation(id uint, userID uint, role string) error
}

type investigationService struct {
	repo                         repositories.InvestigationRepository
	investigationSubtopicService InvestigationSubtopicService
}

func NewInvestigationService(repo repositories.InvestigationRepository, investigationSubtopicService InvestigationSubtopicService) InvestigationService {
	return &investigationService{
		repo:                         repo,
		investigationSubtopicService: investigationSubtopicService,
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

func (s *investigationService) GetAllInvestigationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *investigationService) GetAllInvestigationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *investigationService) CountInvestigationsNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *investigationService) GetInvestigationByID(id uint, userID uint, role string) (*models.Investigation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	investigation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if investigation.IsApproved != nil && *investigation.IsApproved {
		return investigation, nil
	}

	if investigation.UserID == userID || role == "admin" {
		return investigation, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *investigationService) UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error {
	if investigation == nil || investigation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetInvestigationByID(investigation.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	investigation.IsApproved = nil

	updates := utils.StructToMap(investigation)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *investigationService) DeleteInvestigation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetInvestigationByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
