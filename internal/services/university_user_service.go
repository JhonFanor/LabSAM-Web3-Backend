package services

import (
	"fmt"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"gorm.io/gorm"
)

type UniversityUserService interface {
	CreateUniversityUser(universityUser *models.UniversityUser) (*models.UniversityUser, error)
	UpdateUniversityUser(universityUser *models.UniversityUser) error
	DeleteUniversityUser(userID uint) error
	GetUniversityUserByID(userID uint) (*models.UniversityUser, error)
	GetAllUniversityUsers() ([]models.UniversityUser, error)
}

type universityUserService struct {
	repo repositories.UniversityUserRepository
	db   *gorm.DB
}

func NewUniversityUserService(repo repositories.UniversityUserRepository, db *gorm.DB) UniversityUserService {
	return &universityUserService{
		repo: repo,
		db:   db,
	}
}

func (s *universityUserService) CreateUniversityUser(universityUser *models.UniversityUser) (*models.UniversityUser, error) {
	return s.repo.Create(universityUser)
}

func (s *universityUserService) UpdateUniversityUser(universityUser *models.UniversityUser) error {
	if universityUser.UserID == 0 {
		return customerrors.ErrInvalidData
	}

	updates := utils.StructToMap(universityUser)
	fmt.Print(updates)
	if len(updates) == 0 {
		return nil
	}

	existing, err := s.GetUniversityUserByID(universityUser.UserID)
	if err != nil {
		return err
	}
	return s.repo.Update(existing, updates)
}

func (s *universityUserService) DeleteUniversityUser(userID uint) error {
	return s.repo.Delete(userID)
}

func (s *universityUserService) GetUniversityUserByID(userID uint) (*models.UniversityUser, error) {
	return s.repo.GetByUserID(userID)
}

func (s *universityUserService) GetAllUniversityUsers() ([]models.UniversityUser, error) {
	return s.repo.GetAll()
}
