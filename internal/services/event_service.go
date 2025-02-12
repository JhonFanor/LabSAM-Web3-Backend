package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type EventService interface {
	CreateEvent(event *models.Event, userID uint, subtopicIDs []uint) (*models.Event, error)
}

type eventService struct {
	repo                 repositories.EventRepository
	eventSubtopicService EventSubtopicService
	db                   *gorm.DB
	qm                   *gormmanagers.GormQueryManager
}

func NewEventService(repo repositories.EventRepository, eventSubtopicService EventSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) EventService {
	return &eventService{
		repo:                 repo,
		eventSubtopicService: eventSubtopicService,
		db:                   db,
		qm:                   qm,
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
