package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type InvestigationService interface {
	CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error)
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
