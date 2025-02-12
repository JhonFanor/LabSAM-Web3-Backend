package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type JobExchangeSubtopicService interface {
	CreateJobExchangeSubtopic(jobExchangeSubtopic *models.JobExchangeSubtopic) (*models.JobExchangeSubtopic, error)
	DeleteJobExchangeSubtopic(jobExchangeID, subtopicID uint) error
}

type jobExchangeSubtopicService struct {
	repo repositories.JobExchangeSubtopicRepository
	db   *gorm.DB
}

func NewJobExchangeSubtopicService(repo repositories.JobExchangeSubtopicRepository, db *gorm.DB) JobExchangeSubtopicService {
	return &jobExchangeSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *jobExchangeSubtopicService) CreateJobExchangeSubtopic(jobExchangeSubtopic *models.JobExchangeSubtopic) (*models.JobExchangeSubtopic, error) {
	if jobExchangeSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(jobExchangeSubtopic)
}

func (s *jobExchangeSubtopicService) DeleteJobExchangeSubtopic(jobExchangeID, subtopicID uint) error {
	if jobExchangeID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(jobExchangeID, subtopicID)
}
