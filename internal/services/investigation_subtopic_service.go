package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type InvestigationSubtopicService interface {
	CreateInvestigationSubtopic(investigationSubtopic *models.InvestigationSubtopic) (*models.InvestigationSubtopic, error)
	DeleteInvestigationSubtopic(investigationID, subtopicID uint) error
}

type investigationSubtopicService struct {
	repo repositories.InvestigationSubtopicRepository
	db   *gorm.DB
}

func NewInvestigationSubtopicService(repo repositories.InvestigationSubtopicRepository, db *gorm.DB) InvestigationSubtopicService {
	return &investigationSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *investigationSubtopicService) CreateInvestigationSubtopic(investigationSubtopic *models.InvestigationSubtopic) (*models.InvestigationSubtopic, error) {
	if investigationSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(investigationSubtopic)
}

func (s *investigationSubtopicService) DeleteInvestigationSubtopic(investigationID, subtopicID uint) error {
	if investigationID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(investigationID, subtopicID)
}
