package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type BankOfResumeService interface {
	CreateBankOfResume(bankOfResume *models.BankOfResume, userID uint, subtopicIDs []uint) (*models.BankOfResume, error)
	GetAllBankOfResumes(c *gin.Context) (*dto.PaginationDTO, error)
	GetBankOfResumeByID(id uint, userID uint, role string) (*models.BankOfResume, error)
	UpdateBankOfResume(bankOfResume *models.BankOfResume, userID uint, role string) error
	DeleteBankOfResume(id uint, userId uint, role string) error
}

type bankOfResumeService struct {
	repo                        repositories.BankOfResumeRepository
	bankOfResumeSubtopicService BankOfResumeSubtopicService
}

func NewBankOfResumeService(repo repositories.BankOfResumeRepository, bankOfResumeSubtopicService BankOfResumeSubtopicService) BankOfResumeService {
	return &bankOfResumeService{
		repo:                        repo,
		bankOfResumeSubtopicService: bankOfResumeSubtopicService,
	}
}

func (s *bankOfResumeService) CreateBankOfResume(bankOfResume *models.BankOfResume, userID uint, subtopicIDs []uint) (*models.BankOfResume, error) {
	if bankOfResume == nil {
		return nil, customerrors.ErrInvalidData
	}

	bankOfResume.UserID = userID

	createdBankOfResume, err := s.repo.Create(bankOfResume)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		bankOfResumeSubtopic := &models.BankOfResumeSubtopic{
			BankOfResumeID: createdBankOfResume.ID,
			SubtopicID:     uint(subtopicID),
		}

		_, err := s.bankOfResumeSubtopicService.CreateBankOfResumeSubtopic(bankOfResumeSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdBankOfResume, nil
}

func (s *bankOfResumeService) GetAllBankOfResumes(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *bankOfResumeService) GetBankOfResumeByID(id uint, userID uint, role string) (*models.BankOfResume, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	bank, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if bank.IsApproved != nil && *bank.IsApproved {
		return bank, nil
	}

	if userID == bank.UserID || role == "admin" {
		return bank, nil
	}

	return nil, customerrors.ErrForbidden
}

func (s *bankOfResumeService) UpdateBankOfResume(bankOfResume *models.BankOfResume, userID uint, role string) error {
	if bankOfResume == nil || bankOfResume.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetBankOfResumeByID(bankOfResume.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	bankOfResume.IsApproved = nil

	updates := utils.StructToMap(bankOfResume)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *bankOfResumeService) DeleteBankOfResume(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetBankOfResumeByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
