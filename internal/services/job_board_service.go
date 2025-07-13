package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type JobBoardService interface {
	CreateJobBoard(jobBoard *models.JobBoard, userID uint, subtopicIDs []uint) (*models.JobBoard, error)
	GetAllJobsBoard(c *gin.Context) (*dto.PaginationDTO, error)
	GetJobBoardByID(id uint, userID uint, role string) (*models.JobBoard, error)
	UpdateJobBoard(jobBoard *models.JobBoard, userID uint, role string) error
	DeleteJobBoard(id uint, userID uint, role string) error
}

type jobBoardService struct {
	repo                    repositories.JobBoardRepository
	jobBoardSubtopicService JobBoardSubtopicService
}

func NewJobBoardService(repo repositories.JobBoardRepository, jobBoardSubtopicService JobBoardSubtopicService) JobBoardService {
	return &jobBoardService{
		repo:                    repo,
		jobBoardSubtopicService: jobBoardSubtopicService,
	}
}

func (s *jobBoardService) CreateJobBoard(jobBoard *models.JobBoard, userID uint, subtopicIDs []uint) (*models.JobBoard, error) {
	if jobBoard == nil {
		return nil, customerrors.ErrInvalidData
	}

	jobBoard.UserID = userID

	createdJobBoard, err := s.repo.Create(jobBoard)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		jobBoardSubtopic := &models.JobBoardSubtopic{
			JobBoardID: createdJobBoard.ID,
			SubtopicID: uint(subtopicID),
		}

		_, err := s.jobBoardSubtopicService.CreateJobBoardSubtopic(jobBoardSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdJobBoard, nil
}

func (s *jobBoardService) GetAllJobsBoard(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *jobBoardService) GetJobBoardByID(id uint, userID uint, role string) (*models.JobBoard, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if existing.IsApproved != nil && *existing.IsApproved {
		return existing, nil
	}

	if existing.UserID == userID || role == "admin" {
		return existing, nil
	}

	return nil, customerrors.ErrForbidden
}

func (s *jobBoardService) UpdateJobBoard(jobBoard *models.JobBoard, userID uint, role string) error {
	if jobBoard == nil || jobBoard.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetJobBoardByID(jobBoard.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	jobBoard.IsApproved = nil

	updates := utils.StructToMap(jobBoard)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *jobBoardService) DeleteJobBoard(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetJobBoardByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
