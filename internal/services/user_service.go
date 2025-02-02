package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
	db   *gorm.DB
}

func NewUserService(repo repositories.UserRepository, db *gorm.DB) UserService {
	return &userService{
		repo: repo,
		db:   db,
	}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) FindUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) FindUserByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}
