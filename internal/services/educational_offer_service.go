package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type EducationalOfferService interface {
	CreateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, subtopicIDs []uint) (*models.EducationalOffer, error)
	GetAllEducationalOffers(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllEducationalOffersByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllEducationalOffersNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountEducationalOffersNotApproved(role string) (int64, error)
	GetEducationalOfferByID(id uint, userID uint, role string) (*models.EducationalOffer, error)
	UpdateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, role string) error
	SetEducationalOfferApproval(id uint, approved bool, role string) error
	DeleteEducationalOffer(id uint, userID uint, role string) error
}

type educationalOfferService struct {
	repo                            repositories.EducationalOfferRepository
	educationalOfferSubtopicService EducationalOfferSubtopicService
}

func NewEducationalOfferService(repo repositories.EducationalOfferRepository, educationalOfferSubtopicService EducationalOfferSubtopicService) EducationalOfferService {
	return &educationalOfferService{
		repo:                            repo,
		educationalOfferSubtopicService: educationalOfferSubtopicService,
	}
}

func (s *educationalOfferService) CreateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, subtopicIDs []uint) (*models.EducationalOffer, error) {
	if educationalOffer == nil {
		return nil, customerrors.ErrInvalidData
	}

	educationalOffer.UserID = userID

	createdEducationalOffer, err := s.repo.Create(educationalOffer)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		educationalOfferSubtopic := &models.EducationalOfferSubtopic{
			EducationalOfferID: createdEducationalOffer.ID,
			SubtopicID:         uint(subtopicID),
		}

		_, err := s.educationalOfferSubtopicService.CreateEducationalOfferSubtopic(educationalOfferSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdEducationalOffer, nil
}

func (s *educationalOfferService) GetAllEducationalOffers(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *educationalOfferService) GetAllEducationalOffersByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *educationalOfferService) GetAllEducationalOffersNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *educationalOfferService) CountEducationalOffersNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *educationalOfferService) GetEducationalOfferByID(id uint, userID uint, role string) (*models.EducationalOffer, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	educationalOffer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if educationalOffer.IsApproved != nil && *educationalOffer.IsApproved {
		return educationalOffer, nil
	}

	if educationalOffer.UserID == userID || role == "admin" {
		return educationalOffer, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *educationalOfferService) UpdateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, role string) error {
	if educationalOffer == nil || educationalOffer.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetEducationalOfferByID(educationalOffer.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	educationalOffer.IsApproved = nil

	updates := utils.StructToMap(educationalOffer)
	if len(updates) == 0 {
		return nil
	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *educationalOfferService) SetEducationalOfferApproval(id uint, approved bool, role string) error {
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

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *educationalOfferService) DeleteEducationalOffer(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetEducationalOfferByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
