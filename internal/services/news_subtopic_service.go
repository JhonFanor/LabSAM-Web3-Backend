package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type NewsSubtopicService interface {
	CreateNewsSubtopic(newsSubtopic *models.NewsSubtopic) (*models.NewsSubtopic, error)
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
	return s.repo.Create(newsSubtopic)
}

func (s *newsSubtopicService) DeleteNewsSubtopic(newsID, subtopicID uint) error {
	if newsID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(newsID, subtopicID)
}
