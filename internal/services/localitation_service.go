package services

import (
	"errors"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type LocalitationService interface {
	AssignLocalitation(localitation *models.Localitation) (*models.Localitation, error)
	CreateLocalitation(localitation *models.Localitation) (*models.Localitation, error)
	GetLocalitationByFields(localitation *models.Localitation) (*models.Localitation, error)
	DeleteLocalitation(id uint) error
}

type localitationService struct {
	repo repositories.LocalitationRepository
}

func NewLocalitationService(repo repositories.LocalitationRepository) LocalitationService {
	return &localitationService{
		repo: repo,
	}
}

func (s *localitationService) AssignLocalitation(localitation *models.Localitation) (*models.Localitation, error) {
	found, err := s.GetLocalitationByFields(localitation)
	if err == nil {
		return found, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		created, err := s.CreateLocalitation(localitation)
		if err != nil {
			return nil, err
		}
		return created, nil
	}

	return nil, err
}

func (s *localitationService) CreateLocalitation(localitation *models.Localitation) (*models.Localitation, error) {
	if localitation == nil {
		return nil, customerrors.ErrInvalidData
	}

	createdLocalitation, err := s.repo.Create(localitation)
	if err != nil {
		return nil, err
	}

	return createdLocalitation, nil
}

func (s *localitationService) GetLocalitationByFields(localitation *models.Localitation) (*models.Localitation, error) {
	return s.repo.GetByFields(localitation)
}

func (s *localitationService) DeleteLocalitation(id uint) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	return s.repo.Delete(id)
}
