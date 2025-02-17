package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type DocumentationSubtopicService interface {
	CreateDocumentationSubtopic(documentationSubtopic *models.DocumentationSubtopic) (*models.DocumentationSubtopic, error)
	GetDocumentationSubtopicByID(documentationID uint, subtopicID uint) (*models.DocumentationSubtopic, error)
	DeleteDocumentationSubtopic(documentationID, subtopicID uint) error
}

type documentationSubtopicService struct {
	repo repositories.DocumentationSubtopicRepository
	db   *gorm.DB
}

func NewDocumentationSubtopicService(repo repositories.DocumentationSubtopicRepository, db *gorm.DB) DocumentationSubtopicService {
	return &documentationSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *documentationSubtopicService) CreateDocumentationSubtopic(documentationSubtopic *models.DocumentationSubtopic) (*models.DocumentationSubtopic, error) {
	if documentationSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}

	existing, err := s.GetDocumentationSubtopicByID(documentationSubtopic.DocumentationID, documentationSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(documentationSubtopic)
}

func (s *documentationSubtopicService) GetDocumentationSubtopicByID(documentationID uint, subtopicID uint) (*models.DocumentationSubtopic, error) {
	if documentationID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(documentationID, subtopicID)
}

func (s *documentationSubtopicService) DeleteDocumentationSubtopic(documentationID, subtopicID uint) error {
	if documentationID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(documentationID, subtopicID)
}
