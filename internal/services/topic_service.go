package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type TopicService interface {
	CreateTopic(topic *models.Topic) (*models.Topic, error)
	UpdateTopic(topic *models.Topic) error
	DeleteTopic(id uint) error
	GetTopicByID(id uint) (*models.Topic, error)
	GetAllTopic() ([]models.Topic, error)
}

type topicService struct {
	repo repositories.TopicRepository
	db   *gorm.DB
}

func NewTopicService(repo repositories.TopicRepository, db *gorm.DB) TopicService {
	return &topicService{
		repo: repo,
		db:   db,
	}
}

func (s *topicService) CreateTopic(topic *models.Topic) (*models.Topic, error) {
	if topic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(topic)
}

func (s *topicService) UpdateTopic(topic *models.Topic) error {
	if topic == nil || topic.ID == 0 {
		return customerrors.ErrInvalidData
	}
	return s.repo.Update(topic)
}

func (s *topicService) DeleteTopic(id uint) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(id)
}

func (s *topicService) GetTopicByID(id uint) (*models.Topic, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *topicService) GetAllTopic() ([]models.Topic, error) {
	return s.repo.GetAll()
}
