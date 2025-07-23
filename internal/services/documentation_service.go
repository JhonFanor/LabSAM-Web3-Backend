package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/observers"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type DocumentationService interface {
	CreateDocumentation(documentation *models.Documentation, userID uint, subtopicIDs []uint) (*models.Documentation, error)
	GetAllDocumentations(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllDocumentationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllDocumentationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountDocumentationsNotApproved(role string) (int64, error)
	GetDocumentationByID(id uint, userID uint, role string) (*models.Documentation, error)
	UpdateDocumentation(documentation *models.Documentation, userID uint, role string) error
	SetDocumentationApproval(id uint, approved bool, adminId uint, role string) error
	DeleteDocumentation(id uint, userID uint, role string) error
}

type documentationService struct {
	repo                         repositories.DocumentationRepository
	documentationSubtopicService DocumentationSubtopicService
	adminNotificationObserver    *observers.AdminNotificationObserver
	userNotificationObserver     *observers.UserNotificationObserver
}

func NewDocumentationService(repo repositories.DocumentationRepository, documentationSubtopicService DocumentationSubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) DocumentationService {
	return &documentationService{
		repo:                         repo,
		documentationSubtopicService: documentationSubtopicService,
		adminNotificationObserver:    adminNotificationObserver,
		userNotificationObserver:     userNotificationObserver,
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

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado una nueva documentación.",
		Action:       "created",
		ResourceID:   int(createdDocumentation.ID),
		ResourceType: "documentation",
	})

	return createdDocumentation, nil
}

func (s *documentationService) GetAllDocumentations(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *documentationService) GetAllDocumentationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *documentationService) GetAllDocumentationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *documentationService) CountDocumentationsNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *documentationService) GetDocumentationByID(id uint, userID uint, role string) (*models.Documentation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	documentation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if documentation.IsApproved != nil && *documentation.IsApproved {
		return documentation, nil
	}

	if documentation.UserID == userID || role == "admin" {
		return documentation, nil
	}

	return nil, customerrors.ErrInvalidID
}

func (s *documentationService) UpdateDocumentation(documentation *models.Documentation, userID uint, role string) error {
	if documentation == nil || documentation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetDocumentationByID(documentation.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	documentation.IsApproved = nil

	updates := utils.StructToMap(documentation)
	if len(updates) == 0 {
		return nil
	}

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado una documentación.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "documentation",
		})

	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *documentationService) SetDocumentationApproval(id uint, approved bool, adminId uint, role string) error {
	if role != "admin" {
		return customerrors.ErrUnauthorized
	}

	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	isApproved := approved
	updates := map[string]interface{}{
		"is_approved": &isApproved,
	}

	message := ""
	action := ""
	var typeObserver observers.EventObserverType
	if approved {
		typeObserver = observers.EventObserverType(observers.Approved)
		message = "El administrador aprobo la documentación."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo la documentación."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "documentation",
	})

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *documentationService) DeleteDocumentation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetDocumentationByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
