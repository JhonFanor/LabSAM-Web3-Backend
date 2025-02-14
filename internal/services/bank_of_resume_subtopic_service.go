package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type BankOfResumeSubtopicService interface {
	CreateBankOfResumeSubtopic(bankOfResumeSubtopic *models.BankOfResumeSubtopic) (*models.BankOfResumeSubtopic, error)
	DeleteBankOfResumeSubtopic(bankOfResumeID, subtopicID uint) error
}

type bankOfResumeSubtopicService struct {
	repo repositories.BankOfResumeSubtopicRepository
	db   *gorm.DB
}

func NewBankOfResumeSubtopicService(repo repositories.BankOfResumeSubtopicRepository, db *gorm.DB) BankOfResumeSubtopicService {
	return &bankOfResumeSubtopicService{
		repo: repo,
		db:   db,
	}
}

func (s *bankOfResumeSubtopicService) CreateBankOfResumeSubtopic(bankOfResumeSubtopic *models.BankOfResumeSubtopic) (*models.BankOfResumeSubtopic, error) {
	if bankOfResumeSubtopic == nil {
		return nil, customerrors.ErrInvalidData
	}

	existing, err := s.repo.GetByID(bankOfResumeSubtopic.BankOfResumeID, bankOfResumeSubtopic.SubtopicID)

	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return s.repo.Create(bankOfResumeSubtopic)
}

func (s *bankOfResumeSubtopicService) GetBankOfResumeSubtopicByID(bankOfResumeID uint, subtopicID uint) (*models.BankOfResumeSubtopic, error) {
	if bankOfResumeID == 0 || subtopicID == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(bankOfResumeID, subtopicID)
}

func (s *bankOfResumeSubtopicService) DeleteBankOfResumeSubtopic(bankOfResumeID, subtopicID uint) error {
	if bankOfResumeID == 0 || subtopicID == 0 {
		return customerrors.ErrInvalidID
	}
	return s.repo.Delete(bankOfResumeID, subtopicID)
}
