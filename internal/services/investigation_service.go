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

type InvestigationService interface {
	CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error)
	GetAllInvestigations(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllInvestigationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllInvestigationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountInvestigationsNotApproved(role string) (int64, error)
	GetInvestigationByID(id uint, userID uint, role string) (*models.Investigation, error)
	UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error
	SetInvestigationApproval(id uint, approved bool, adminId uint, role string) error
	DeleteInvestigation(id uint, userID uint, role string) error
}

type investigationService struct {
	repo                         repositories.InvestigationRepository
	investigationSubtopicService InvestigationSubtopicService
	adminNotificationObserver    *observers.AdminNotificationObserver
	userNotificationObserver     *observers.UserNotificationObserver
}

func NewInvestigationService(repo repositories.InvestigationRepository, investigationSubtopicService InvestigationSubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) InvestigationService {
	return &investigationService{
		repo:                         repo,
		investigationSubtopicService: investigationSubtopicService,
		adminNotificationObserver:    adminNotificationObserver,
		userNotificationObserver:     userNotificationObserver,
	}
}

func (s *investigationService) CreateInvestigation(investigation *models.Investigation, userID uint, subtopicIDs []uint) (*models.Investigation, error) {
	if investigation == nil {
		return nil, customerrors.ErrInvalidData
	}

	investigation.UserID = userID

	createdInvestigation, err := s.repo.Create(investigation)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		investigationSubtopic := &models.InvestigationSubtopic{
			InvestigationID: createdInvestigation.ID,
			SubtopicID:      uint(subtopicID),
		}

		_, err := s.investigationSubtopicService.CreateInvestigationSubtopic(investigationSubtopic)
		if err != nil {
			return nil, err
		}
	}

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado una nueva investigación.",
		Action:       "created",
		ResourceID:   int(createdInvestigation.ID),
		ResourceType: "investigation",
	})

	return createdInvestigation, nil
}

func (s *investigationService) GetAllInvestigations(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *investigationService) GetAllInvestigationsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *investigationService) GetAllInvestigationsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *investigationService) CountInvestigationsNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *investigationService) GetInvestigationByID(id uint, userID uint, role string) (*models.Investigation, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	investigation, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if investigation.IsApproved != nil && *investigation.IsApproved {
		return investigation, nil
	}

	if investigation.UserID == userID || role == "admin" {
		return investigation, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *investigationService) UpdateInvestigation(investigation *models.Investigation, userID uint, role string) error {
	if investigation == nil || investigation.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetInvestigationByID(investigation.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	investigation.IsApproved = nil

	updates := utils.StructToMap(investigation)
	if len(updates) == 0 {
		return nil
	}

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado una hoja de vida.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "bank_of_resume",
		})

	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *investigationService) SetInvestigationApproval(id uint, approved bool, adminId uint, role string) error {
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
		message = "El administrador aprobo tú hoja de vida."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo tú hoja de vida."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "bank_of_resume",
	})

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *investigationService) DeleteInvestigation(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetInvestigationByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
