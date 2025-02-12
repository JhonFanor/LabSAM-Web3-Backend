package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type DocumentationService interface {
	CreateDocumentation(documentation *models.Documentation, userID uint, subtopicIDs []uint) (*models.Documentation, error)
}

type documentationService struct {
	repo                         repositories.DocumentationRepository
	documentationSubtopicService DocumentationSubtopicService
	db                           *gorm.DB
	qm                           *gormmanagers.GormQueryManager
}

func NewDocumentationService(repo repositories.DocumentationRepository, documentationSubtopicService DocumentationSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) DocumentationService {
	return &documentationService{
		repo:                         repo,
		documentationSubtopicService: documentationSubtopicService,
		db:                           db,
		qm:                           qm,
	}
}

func (s *documentationService) CreateDocumentation(documentation *models.Documentation, userID uint, subtopicIDs []uint) (*models.Documentation, error) {
	if documentation == nil {
		return nil, customerrors.ErrInvalidData
	}
	documentation.UserID = userID

	createdDocumentation, err := s.repo.Create(documentation)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		documentationSubtopic := &models.DocumentationSubtopic{
			DocumentationID: createdDocumentation.ID,
			SubtopicID:      uint(subtopicID),
		}

		_, err := s.documentationSubtopicService.CreateDocumentationSubtopic(documentationSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdDocumentation, nil
}
