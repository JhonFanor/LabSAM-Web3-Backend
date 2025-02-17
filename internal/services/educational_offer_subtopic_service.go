package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type EducationalOfferSubtopicService interface {
	CreateEducationalOfferSubtopic(educationalOfferSubtopic *models.EducationalOfferSubtopic) (*models.EducationalOfferSubtopic, error)
	GetEducationalOfferByID(educationalOfferID uint, subtopicID uint) (*models.EducationalOfferSubtopic, error)
	DeleteEducationalOfferSubtopic(educationalOfferID, subtopicID uint) error
}

type educationalOfferSubtopicService struct {
	repo repositories.EducationalOfferSubtopicRepository
	db   *gorm.DB
}

func NewEducationalOfferSubtopicService(repo repositories.EducationalOfferSubtopicRepository, db *gorm.DB) EducationalOfferSubtopicService {
	return &educationalOfferSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *educationalOfferSubtopicService) CreateEducationalOfferSubtopic(educationalOfferSubtopic *models.EducationalOfferSubtopic) (*models.EducationalOfferSubtopic, error) {
	if educationalOfferSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}

	existing, err := s.GetEducationalOfferByID(educationalOfferSubtopic.EducationalOfferID, educationalOfferSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(educationalOfferSubtopic)
}

func (s *educationalOfferSubtopicService) GetEducationalOfferByID(educationalOfferID uint, subtopicID uint) (*models.EducationalOfferSubtopic, error) {
	if educationalOfferID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(educationalOfferID, subtopicID)
}

func (s *educationalOfferSubtopicService) DeleteEducationalOfferSubtopic(educationalOfferID, subtopicID uint) error {
	if educationalOfferID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(educationalOfferID, subtopicID)
}
