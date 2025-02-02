package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type BusinessUserService interface {
	CreateBusinessUser(businessUser *models.BusinessUser) (*models.BusinessUser, error)
	UpdateBusinessUser(businessUser *models.BusinessUser) error
	DeleteBusinessUser(userID uint) error
	GetBusinessUserByUserID(userID uint) (*models.BusinessUser, error)
	GetAllBusinessUsers() ([]models.BusinessUser, error)
}

type businessUserService struct {
	repo repositories.BusinessUserRepository
	db   *gorm.DB
}

func NewBusinessUserService(repo repositories.BusinessUserRepository, db *gorm.DB) BusinessUserService {
	return &businessUserService{
		repo: repo,
		db:   db,
	}
}

func (s *businessUserService) CreateBusinessUser(businessUser *models.BusinessUser) (*models.BusinessUser, error) {
	return s.repo.Create(businessUser)
}

func (s *businessUserService) UpdateBusinessUser(businessUser *models.BusinessUser) error {
	return s.repo.Update(businessUser)
}

func (s *businessUserService) DeleteBusinessUser(userID uint) error {
	return s.repo.Delete(userID)
}

func (s *businessUserService) GetBusinessUserByUserID(userID uint) (*models.BusinessUser, error) {
	return s.repo.GetByUserID(userID)
}

func (s *businessUserService) GetAllBusinessUsers() ([]models.BusinessUser, error) {
	return s.repo.GetAll()
}
