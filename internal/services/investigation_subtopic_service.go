package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type InvestigationSubtopicService interface {
	CreateInvestigationSubtopic(investigationSubtopic *models.InvestigationSubtopic) (*models.InvestigationSubtopic, error)
	GetInvestigationSubtopicByID(investigationID uint, subtopicID uint) (*models.InvestigationSubtopic, error)
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

	existing, err := s.GetInvestigationSubtopicByID(investigationSubtopic.InvestigationID, investigationSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(investigationSubtopic)
}

func (s *investigationSubtopicService) GetInvestigationSubtopicByID(investigationID uint, subtopicID uint) (*models.InvestigationSubtopic, error) {
	if investigationID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(investigationID, subtopicID)
}

func (s *investigationSubtopicService) DeleteInvestigationSubtopic(investigationID, subtopicID uint) error {
	if investigationID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(investigationID, subtopicID)
}
