package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type RegularUserService interface {
	CreateRegularUser(regularUser *models.RegularUser) (*models.RegularUser, error)
	UpdateRegularUser(regularUser *models.RegularUser) error
	DeleteRegularUser(userID uint) error
	GetRegularUserByID(userID uint) (*models.RegularUser, error)
	GetAllRegularUsers() ([]models.RegularUser, error)
}

type regularUserService struct {
	repo repositories.RegularUserRepository
	db   *gorm.DB
}

func NewRegularUserService(repo repositories.RegularUserRepository, db *gorm.DB) RegularUserService {
	return &regularUserService{
		repo: repo,
		db:   db,
	}
}

func (s *regularUserService) CreateRegularUser(regularUser *models.RegularUser) (*models.RegularUser, error) {
	return s.repo.Create(regularUser)
}

func (s *regularUserService) UpdateRegularUser(regularUser *models.RegularUser) error {
	return s.repo.Update(regularUser)
}

func (s *regularUserService) DeleteRegularUser(userID uint) error {
	return s.repo.Delete(userID)
}

func (s *regularUserService) GetRegularUserByID(userID uint) (*models.RegularUser, error) {
	return s.repo.GetByUserID(userID)
}

func (s *regularUserService) GetAllRegularUsers() ([]models.RegularUser, error) {
	return s.repo.GetAll()
}
