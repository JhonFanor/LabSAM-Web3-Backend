package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type EventService interface {
	CreateEvent(event *models.Event, userID uint, subtopicIDs []uint) (*models.Event, error)
	GetAllEvents(c *gin.Context) (*dto.PaginationDTO, error)
	GetEventByID(id uint) (*models.Event, error)
	UpdateEvent(event *models.Event, userID uint, role string) error
	DeleteEvent(id uint, userID uint, role string) error
}

type eventService struct {
	repo                 repositories.EventRepository
	eventSubtopicService EventSubtopicService
}

func NewEventService(repo repositories.EventRepository, eventSubtopicService EventSubtopicService) EventService {
	return &eventService{
		repo:                 repo,
		eventSubtopicService: eventSubtopicService,
	}
}

func (s *eventService) CreateEvent(event *models.Event, userID uint, subtopicIDs []uint) (*models.Event, error) {
	if event == nil {
		return nil, customerrors.ErrInvalidData
	}

	event.UserID = userID

	createdEvent, err := s.repo.Create(event)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		eventSubtopic := &models.EventSubtopic{
			EventID:    createdEvent.ID,
			SubtopicID: uint(subtopicID),
		}

		_, err := s.eventSubtopicService.CreateEventSubtopic(eventSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdEvent, nil
}

func (s *eventService) GetAllEvents(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *eventService) GetEventByID(id uint) (*models.Event, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *eventService) UpdateEvent(event *models.Event, userID uint, role string) error {
	if event == nil || event.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetEventByID(event.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(event)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *eventService) DeleteEvent(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetEventByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
