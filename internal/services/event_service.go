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

type EventService interface {
	CreateEvent(event *models.Event, userID uint, subtopicIDs []uint) (*models.Event, error)
	GetAllEvents(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllEventsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllEventsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountEventsNotApproved(role string) (int64, error)
	GetEventByID(id uint, userID uint, role string) (*models.Event, error)
	UpdateEvent(event *models.Event, userID uint, role string) error
	SetEventApproval(id uint, approved bool, adminId uint, role string) error
	DeleteEvent(id uint, userID uint, role string) error
}

type eventService struct {
	repo                      repositories.EventRepository
	localitationService       LocalitationService
	eventSubtopicService      EventSubtopicService
	adminNotificationObserver *observers.AdminNotificationObserver
	userNotificationObserver  *observers.UserNotificationObserver
}

func NewEventService(repo repositories.EventRepository, eventSubtopicService EventSubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) EventService {
	return &eventService{
		repo:                      repo,
		eventSubtopicService:      eventSubtopicService,
		adminNotificationObserver: adminNotificationObserver,
		userNotificationObserver:  userNotificationObserver,
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

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado un nuevo evento.",
		Action:       "created",
		ResourceID:   int(createdEvent.ID),
		ResourceType: "event",
	})

	return createdEvent, nil
}

func (s *eventService) GetAllEvents(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *eventService) GetAllEventsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *eventService) GetAllEventsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *eventService) CountEventsNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
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

	if event.UserID == id || role == "admin" {
		return event, nil
	}

	return nil, customerrors.ErrUnauthorized
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

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado un evento.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "event",
		})

	}

	existing.User = nil
	err = s.repo.Update(existing, updates)

	if err == nil && event.LocalitationID != nil && *event.LocalitationID != 0 && event.LocalitationID != existing.LocalitationID {
		s.localitationService.DeleteLocalitation(*existing.LocalitationID)
	}

	return err
}

func (s *eventService) SetEventApproval(id uint, approved bool, adminId uint, role string) error {
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
		message = "El administrador aprobo tú evento."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo tú evento."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "event",
	})

	existing.User = nil
	return s.repo.Update(existing, updates)
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
