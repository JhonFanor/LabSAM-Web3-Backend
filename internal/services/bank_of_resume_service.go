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

type BankOfResumeService interface {
	CreateBankOfResume(bankOfResume *models.BankOfResume, userID uint, subtopicIDs []uint) (*models.BankOfResume, error)
	UpdateBankOfResume(bankOfResume *models.BankOfResume, userID uint, role string) error
	DeleteBankOfResume(id uint, userId uint, role string) error
	GetBankOfResumeByID(id uint) (*models.BankOfResume, error)
	GetAllBankOfResumes(c *gin.Context) (*dto.PaginationDTO, error)
}

type bankOfResumeService struct {
	repo                        repositories.BankOfResumeRepository
	bankOfResumeSubtopicService BankOfResumeSubtopicService
	db                          *gorm.DB
	qm                          *gormmanagers.GormQueryManager
}

func NewBankOfResumeService(repo repositories.BankOfResumeRepository, bankOfResumeSubtopicService BankOfResumeSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) BankOfResumeService {
	return &bankOfResumeService{
		repo:                        repo,
		bankOfResumeSubtopicService: bankOfResumeSubtopicService,
		db:                          db,
		qm:                          qm,
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

func (s *bankOfResumeService) UpdateBankOfResume(bankOfResume *models.BankOfResume, userID uint, role string) error {
	if bankOfResume == nil || bankOfResume.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetBankOfResumeByID(bankOfResume.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.GetModifiedFields(existing, bankOfResume)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(bankOfResume)
}

func (s *bankOfResumeService) DeleteBankOfResume(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetBankOfResumeByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}

func (s *bankOfResumeService) GetBankOfResumeByID(id uint) (*models.BankOfResume, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *bankOfResumeService) GetAllBankOfResumes(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}
