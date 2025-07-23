package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type SubtopicService interface {
	CreateSubtopic(topic *models.Subtopic) (*models.Subtopic, error)
	UpdateSubtopic(topic *models.Subtopic) error
	DeleteSubtopic(id uint) error
	GetSubtopicByID(id uint) (*models.Subtopic, error)
	GetAllSubtopic() ([]models.Subtopic, error)
}

type subtopicService struct {
	repo repositories.SubtopicRepository
	db   *gorm.DB
}

func NewSubtopicService(repo repositories.SubtopicRepository, db *gorm.DB) SubtopicService {
	return &subtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *subtopicService) CreateSubtopic(subtopic *models.Subtopic) (*models.Subtopic, error) {
	if subtopic == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(subtopic)
}

func (s *subtopicService) UpdateSubtopic(subtopic *models.Subtopic) error {
	if subtopic == nil || subtopic.ID == 0 {
		return customerrors.ErrInvalidData
	}
	return s.repo.Update(subtopic)
}

func (s *subtopicService) DeleteSubtopic(id uint) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(id)
}

func (s *subtopicService) GetSubtopicByID(id uint) (*models.Subtopic, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *subtopicService) GetAllSubtopic() ([]models.Subtopic, error) {
	return s.repo.GetAll()
}
