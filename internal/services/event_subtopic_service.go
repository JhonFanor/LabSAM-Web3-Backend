package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type EventSubtopicService interface {
	CreateEventSubtopic(eventSubtopic *models.EventSubtopic) (*models.EventSubtopic, error)
	DeleteEventSubtopic(eventID, subtopicID uint) error
}

type eventSubtopicService struct {
	repo repositories.EventSubtopicRepository
	db   *gorm.DB
}

func NewEventSubtopicService(repo repositories.EventSubtopicRepository, db *gorm.DB) EventSubtopicService {
	return &eventSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *eventSubtopicService) CreateEventSubtopic(eventSubtopic *models.EventSubtopic) (*models.EventSubtopic, error) {
	if eventSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(eventSubtopic)
}

func (s *eventSubtopicService) DeleteEventSubtopic(eventID, subtopicID uint) error {
	if eventID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(eventID, subtopicID)
}
