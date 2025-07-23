package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/observers"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type CompanyService interface {
	CreateCompany(company *models.Company, userID uint, subtopicIDs []uint) (*models.Company, error)
	GetAllCompanies(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllCompaniesByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllCompaniesNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountCompaniesNotApproved(role string) (int64, error)
	CountCompanyBySubtopic() ([]responses.SubtopicCountResponse, error)
	GetCompanyByID(id uint, userID uint, role string) (*models.Company, error)
	UpdateCompany(company *models.Company, userID uint, role string) error
	SetCompanyApproval(id uint, approved bool, adminId uint, role string) error
	DeleteCompany(id uint, userID uint, role string) error
}

type companyService struct {
	repo                      repositories.CompanyRepository
	localitationService       LocalitationService
	companySubtopicService    CompanySubtopicService
	subtopicService           SubtopicService
	adminNotificationObserver *observers.AdminNotificationObserver
	userNotificationObserver  *observers.UserNotificationObserver
}

func NewCompanyService(repo repositories.CompanyRepository, localitationService LocalitationService, companySubtopicService CompanySubtopicService, subtopicService SubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) CompanyService {
	return &companyService{
		repo:                      repo,
		localitationService:       localitationService,
		companySubtopicService:    companySubtopicService,
		subtopicService:           subtopicService,
		adminNotificationObserver: adminNotificationObserver,
		userNotificationObserver:  userNotificationObserver,
	}
}

func (s *companyService) CreateCompany(company *models.Company, userID uint, subtopicIDs []uint) (*models.Company, error) {
	if company == nil {
		return nil, customerrors.ErrInvalidData
	}

	company.UserID = userID

	createdCompany, err := s.repo.Create(company)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		companySubtopic := &models.CompanySubtopic{
			CompanyID:  createdCompany.ID,
			SubtopicID: uint(subtopicID),
		}

		_, err := s.companySubtopicService.CreateCompanySubtopic(companySubtopic)
		if err != nil {
			return nil, err
		}
	}

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado información de una empresa.",
		Action:       "created",
		ResourceID:   int(createdCompany.ID),
		ResourceType: "company",
	})

	return createdCompany, nil
}

func (s *companyService) GetAllCompanies(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *companyService) GetAllCompaniesByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *companyService) GetAllCompaniesNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *companyService) CountCompaniesNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *companyService) CountCompanyBySubtopic() ([]responses.SubtopicCountResponse, error) {
	subtopics, err := s.subtopicService.GetAllSubtopic()
	if err != nil {
		return nil, err
	}

	var result []responses.SubtopicCountResponse

	for _, sub := range subtopics {
		count, err := s.repo.CountBySubtopicID(sub.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, responses.SubtopicCountResponse{
			SubtopicName: sub.Name,
			Count:        count,
		})
	}

	return result, nil
}

func (s *companyService) GetCompanyByID(id uint, userID uint, role string) (*models.Company, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	company, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if company.IsApproved != nil && *company.IsApproved {
		return company, nil
	}

	if company.UserID == userID || role == "admin" {
		return company, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *companyService) UpdateCompany(company *models.Company, userID uint, role string) error {
	if company == nil || company.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetCompanyByID(company.ID, userID, role)
	if err != nil {
		return err
	}

	company.IsApproved = nil

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(company)
	if len(updates) == 0 {
		return nil
	}

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado la información de una empresa.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "company",
		})

	}

	existing.User = nil
	err = s.repo.Update(existing, updates)

	if err == nil && company.LocalitationID != nil && existing.LocalitationID != nil && *company.LocalitationID != *existing.LocalitationID {
		s.localitationService.DeleteLocalitation(*existing.LocalitationID)
	}

	return err
}

func (s *companyService) SetCompanyApproval(id uint, approved bool, adminId uint, role string) error {
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
		message = "El administrador aprobo la información de la empresa."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo la información de la empresa."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "company",
	})

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *companyService) DeleteCompany(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetCompanyByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
