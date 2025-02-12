package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type LegislationService interface {
	CreateLegislation(legislation *models.Legislation, userID uint, subtopicIDs []uint) (*models.Legislation, error)
}

type legislationService struct {
	repo                       repositories.LegislationRepository
	legislationSubtopicService LegislationSubtopicService
	db                         *gorm.DB
	qm                         *gormmanagers.GormQueryManager
}

func NewLegislationService(repo repositories.LegislationRepository, legislationSubtopicService LegislationSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) LegislationService {
	return &legislationService{
		repo:                       repo,
		legislationSubtopicService: legislationSubtopicService,
		db:                         db,
		qm:                         qm,
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
