package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentationService interface {
	CreateDocumentation(documentation *models.Documentation, userID uint, subtopicIDs []uint) (*models.Documentation, error)
	GetAllDocumentations(c *gin.Context) (*dto.PaginationDTO, error)
	GetDocumentationByID(id uint) (*models.Documentation, error)
	UpdateDocumentation(documentation *models.Documentation, userID uint, role string) error
	DeleteDocumentation(id uint, userID uint, role string) error
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

func (s *documentationService) GetAllDocumentations(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *documentationService) GetDocumentationByID(id uint) (*models.Documentation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *documentationService) UpdateDocumentation(documentation *models.Documentation, userID uint, role string) error {
	if documentation == nil || documentation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetDocumentationByID(documentation.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.GetModifiedFields(existing, documentation)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(documentation)
}

func (s *documentationService) DeleteDocumentation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetDocumentationByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
