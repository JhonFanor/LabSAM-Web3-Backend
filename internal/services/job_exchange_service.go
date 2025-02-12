package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type JobExchangeService interface {
	CreateJobExchange(jobExchange *models.JobExchange, userID uint, subtopicIDs []uint) (*models.JobExchange, error)
}

type jobExchangeService struct {
	repo                       repositories.JobExchangeRepository
	jobExchangeSubtopicService JobExchangeSubtopicService
	db                         *gorm.DB
	qm                         *gormmanagers.GormQueryManager
}

func NewJobExchangeService(repo repositories.JobExchangeRepository, jobExchangeSubtopicService JobExchangeSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) JobExchangeService {
	return &jobExchangeService{
		repo:                       repo,
		jobExchangeSubtopicService: jobExchangeSubtopicService,
		db:                         db,
		qm:                         qm,
	}
}

func (s *jobExchangeService) CreateJobExchange(jobExchange *models.JobExchange, userID uint, subtopicIDs []uint) (*models.JobExchange, error) {
	if jobExchange == nil {
		return nil, customerrors.ErrInvalidData
	}

	jobExchange.UserID = userID

	createdJobExchange, err := s.repo.Create(jobExchange)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		jobExchangeSubtopic := &models.JobExchangeSubtopic{
			JobExchangeID: createdJobExchange.ID,
			SubtopicID:    uint(subtopicID),
		}

		_, err := s.jobExchangeSubtopicService.CreateJobExchangeSubtopic(jobExchangeSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdJobExchange, nil
}
