package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type LegislationService interface {
	CreateLegislation(legislation *models.Legislation, userID uint, subtopicIDs []uint) (*models.Legislation, error)
	GetAllLegislations(c *gin.Context) (*dto.PaginationDTO, error)
	GetLegislationByID(id uint, userID uint, role string) (*models.Legislation, error)
	UpdateLegislation(legislation *models.Legislation, userID uint, role string) error
	DeleteLegislation(id uint, userID uint, role string) error
}

type legislationService struct {
	repo                       repositories.LegislationRepository
	legislationSubtopicService LegislationSubtopicService
}

func NewLegislationService(repo repositories.LegislationRepository, legislationSubtopicService LegislationSubtopicService) LegislationService {
	return &legislationService{
		repo:                       repo,
		legislationSubtopicService: legislationSubtopicService,
	}
}

func (s *legislationService) CreateLegislation(legislation *models.Legislation, userID uint, subtopicIDs []uint) (*models.Legislation, error) {
	if legislation == nil {
		return nil, customerrors.ErrInvalidData
	}

	legislation.UserID = userID

	createdLegislation, err := s.repo.Create(legislation)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		legislationSubtopic := &models.LegislationSubtopic{
			LegislationID: createdLegislation.ID,
			SubtopicID:    uint(subtopicID),
		}

		_, err := s.legislationSubtopicService.CreateLegislationSubtopic(legislationSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdLegislation, nil
}

func (s *legislationService) GetAllLegislations(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *legislationService) GetLegislationByID(id uint, userID uint, role string) (*models.Legislation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	legislation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if legislation.IsApproved {
		return legislation, nil
	}

	if userID == legislation.UserID || role == "admin" {
		return legislation, nil
	}

	return nil, customerrors.ErrForbidden
}

func (s *legislationService) UpdateLegislation(legislation *models.Legislation, userID uint, role string) error {
	if legislation == nil || legislation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetLegislationByID(legislation.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(legislation)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *legislationService) DeleteLegislation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetLegislationByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
