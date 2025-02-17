package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type NewsSubtopicService interface {
	CreateNewsSubtopic(newsSubtopic *models.NewsSubtopic) (*models.NewsSubtopic, error)
	GetNewsSubtopicByID(newsID, subtopicID uint) (*models.NewsSubtopic, error)
	DeleteNewsSubtopic(newsID, subtopicID uint) error
}

type newsSubtopicService struct {
	repo repositories.NewsSubtopicRepository
	db   *gorm.DB
}

func NewNewsSubtopicService(repo repositories.NewsSubtopicRepository, db *gorm.DB) NewsSubtopicService {
	return &newsSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *newsSubtopicService) CreateNewsSubtopic(newsSubtopic *models.NewsSubtopic) (*models.NewsSubtopic, error) {
	if newsSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}

	existing, err := s.GetNewsSubtopicByID(newsSubtopic.NewsID, newsSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(newsSubtopic)
}

func (s *newsSubtopicService) GetNewsSubtopicByID(newsID uint, subtopicID uint) (*models.NewsSubtopic, error) {
	if newsID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(newsID, subtopicID)
}

func (s *newsSubtopicService) DeleteNewsSubtopic(newsID, subtopicID uint) error {
	if newsID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(newsID, subtopicID)
}
