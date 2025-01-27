package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type UserService interface {
	FindUserByEmail(email string) (*models.User, error)       // Método exportado
	FindUserByUsername(username string) (*models.User, error) // Método exportado
}

type userService struct {
	userRepo repositories.UserRepository
	db       *gorm.DB
}

func NewUserService(userRepo repositories.UserRepository, db *gorm.DB) UserService {
	return &userService{
		userRepo: userRepo,
		db:       db,
	}
}

func (s *userService) FindUserByEmail(email string) (*models.User, error) { // Método exportado
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) FindUserByUsername(username string) (*models.User, error) { // Método exportado
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
