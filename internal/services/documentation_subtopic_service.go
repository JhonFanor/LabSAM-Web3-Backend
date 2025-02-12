package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type DocumentationSubtopicService interface {
	CreateDocumentationSubtopic(documentationSubtopic *models.DocumentationSubtopic) (*models.DocumentationSubtopic, error)
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
	return s.repo.Create(documentationSubtopic)
}

func (s *documentationSubtopicService) DeleteDocumentationSubtopic(documentationID, subtopicID uint) error {
	if documentationID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(documentationID, subtopicID)
}
