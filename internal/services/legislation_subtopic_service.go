package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type LegislationSubtopicService interface {
	CreateLegislationSubtopic(legislationSubtopic *models.LegislationSubtopic) (*models.LegislationSubtopic, error)
	DeleteLegislationSubtopic(legislationID, subtopicID uint) error
}

type legislationSubtopicService struct {
	repo repositories.LegislationSubtopicRepository
	db   *gorm.DB
}

func NewLegislationSubtopicService(repo repositories.LegislationSubtopicRepository, db *gorm.DB) LegislationSubtopicService {
	return &legislationSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *legislationSubtopicService) CreateLegislationSubtopic(legislationSubtopic *models.LegislationSubtopic) (*models.LegislationSubtopic, error) {
	if legislationSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(legislationSubtopic)
}

func (s *legislationSubtopicService) DeleteLegislationSubtopic(legislationID, subtopicID uint) error {
	if legislationID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(legislationID, subtopicID)
}
