package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type JobBoardSubtopicService interface {
	CreateJobBoardSubtopic(jobBoardSubtopic *models.JobBoardSubtopic) (*models.JobBoardSubtopic, error)
	GetJobBoardSubtopicByID(jobBoardID uint, subtopicID uint) (*models.JobBoardSubtopic, error)
	DeleteJobBoardSubtopic(jobBoardID, subtopicID uint) error
}

type jobBoardSubtopicService struct {
	repo repositories.JobBoardSubtopicRepository
	db   *gorm.DB
}

func NewJobBoardSubtopicService(repo repositories.JobBoardSubtopicRepository, db *gorm.DB) JobBoardSubtopicService {
	return &jobBoardSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *jobBoardSubtopicService) CreateJobBoardSubtopic(jobBoardSubtopic *models.JobBoardSubtopic) (*models.JobBoardSubtopic, error) {
	if jobBoardSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}

	existing, err := s.GetJobBoardSubtopicByID(jobBoardSubtopic.JobBoardID, jobBoardSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(jobBoardSubtopic)
}

func (s *jobBoardSubtopicService) GetJobBoardSubtopicByID(jobBoardID uint, subtopicID uint) (*models.JobBoardSubtopic, error) {
	if jobBoardID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}

	return s.repo.GetByID(jobBoardID, subtopicID)
}

func (s *jobBoardSubtopicService) DeleteJobBoardSubtopic(jobBoardID, subtopicID uint) error {
	if jobBoardID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(jobBoardID, subtopicID)
}
