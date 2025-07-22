package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/observers"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type JobBoardService interface {
	CreateJobBoard(jobBoard *models.JobBoard, userID uint, subtopicIDs []uint) (*models.JobBoard, error)
	GetAllJobsBoard(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllJobsBoardByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllJobsBoardNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountJobsBoardNotApproved(role string) (int64, error)
	GetJobBoardByID(id uint, userID uint, role string) (*models.JobBoard, error)
	UpdateJobBoard(jobBoard *models.JobBoard, userID uint, role string) error
	SetJobBoardApproval(id uint, approved bool, adminId uint, role string) error
	DeleteJobBoard(id uint, userID uint, role string) error
}

type jobBoardService struct {
	repo                      repositories.JobBoardRepository
	jobBoardSubtopicService   JobBoardSubtopicService
	adminNotificationObserver *observers.AdminNotificationObserver
	userNotificationObserver  *observers.UserNotificationObserver
}

func NewJobBoardService(repo repositories.JobBoardRepository, jobBoardSubtopicService JobBoardSubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) JobBoardService {
	return &jobBoardService{
		repo:                      repo,
		jobBoardSubtopicService:   jobBoardSubtopicService,
		adminNotificationObserver: adminNotificationObserver,
		userNotificationObserver:  userNotificationObserver,
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

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado una nueva oferta de trabajo.",
		Action:       "created",
		ResourceID:   int(createdJobBoard.ID),
		ResourceType: "job_board",
	})

	return createdJobBoard, nil
}

func (s *jobBoardService) GetAllJobsBoard(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *jobBoardService) GetAllJobsBoardByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *jobBoardService) GetAllJobsBoardNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *jobBoardService) CountJobsBoardNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *jobBoardService) GetJobBoardByID(id uint, userID uint, role string) (*models.JobBoard, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	jobBoard, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if jobBoard.IsApproved != nil && *jobBoard.IsApproved {
		return jobBoard, nil
	}

	if jobBoard.UserID == userID || role == "admin" {
		return jobBoard, nil
	}

	return nil, customerrors.ErrUnauthorized
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

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado una oferta de trabajo.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "job_board",
		})

	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *jobBoardService) SetJobBoardApproval(id uint, approved bool, adminId uint, role string) error {
	if role != "admin" {
		return customerrors.ErrUnauthorized
	}

	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	isApproved := approved
	updates := map[string]interface{}{
		"is_approved": &isApproved,
	}

	message := ""
	action := ""
	var typeObserver observers.EventObserverType
	if approved {
		typeObserver = observers.EventObserverType(observers.Approved)
		message = "El administrador aprobo tú oferta de trabajo."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo tú oferta de trabajo."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "job_board",
	})

	existing.User = nil
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
