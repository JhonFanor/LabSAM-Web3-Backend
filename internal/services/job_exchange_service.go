package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type JobExchangeService interface {
	CreateJobExchange(jobExchange *models.JobExchange, userID uint, subtopicIDs []uint) (*models.JobExchange, error)
	GetAllJobsExchange(c *gin.Context) (*dto.PaginationDTO, error)
	GetJobExchangeByID(id uint) (*models.JobExchange, error)
	UpdateJobExchange(jobExchange *models.JobExchange, userID uint, role string) error
	DeleteJobExchange(id uint, userID uint, role string) error
}

type jobExchangeService struct {
	repo                       repositories.JobExchangeRepository
	jobExchangeSubtopicService JobExchangeSubtopicService
}

func NewJobExchangeService(repo repositories.JobExchangeRepository, jobExchangeSubtopicService JobExchangeSubtopicService) JobExchangeService {
	return &jobExchangeService{
		repo:                       repo,
		jobExchangeSubtopicService: jobExchangeSubtopicService,
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

func (s *jobExchangeService) GetAllJobsExchange(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *jobExchangeService) GetJobExchangeByID(id uint) (*models.JobExchange, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *jobExchangeService) UpdateJobExchange(jobExchange *models.JobExchange, userID uint, role string) error {
	if jobExchange == nil || jobExchange.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetJobExchangeByID(jobExchange.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(jobExchange)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *jobExchangeService) DeleteJobExchange(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetJobExchangeByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
