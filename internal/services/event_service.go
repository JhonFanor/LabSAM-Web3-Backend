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
	GetEventByID(id uint, userID uint, role string) (*models.Event, error)
	UpdateEvent(event *models.Event, userID uint, role string) error
	DeleteEvent(id uint, userID uint, role string) error
}

type eventService struct {
	repo                 repositories.EventRepository
	localitationService  LocalitationService
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

func (s *eventService) GetEventByID(id uint, userID uint, role string) (*models.Event, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	event, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if event.IsApproved != nil && *event.IsApproved {
		return event, nil
	}

	if event.UserID == 0 || event.UserID == id || role == "admin" {
		return event, nil
	}

	return nil, customerrors.ErrForbidden
}

func (s *eventService) UpdateEvent(event *models.Event, userID uint, role string) error {
	if event == nil || event.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetEventByID(event.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	event.IsApproved = nil

	updates := utils.StructToMap(event)
	if len(updates) == 0 {
		return nil
	}

	err = s.repo.Update(existing, updates)

	if err == nil && *event.LocalitationID != 0 && event.LocalitationID != existing.LocalitationID {
		s.localitationService.DeleteLocalitation(*existing.LocalitationID)
	}

	return err
}

func (s *eventService) DeleteEvent(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetEventByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
