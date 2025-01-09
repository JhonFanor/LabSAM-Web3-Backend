package services

import (
	"errors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/pkg/security"

	"gorm.io/gorm"
)

type AuthService interface {
	RegisterUser(user *models.User) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
	db       *gorm.DB
}

func NewAuthService(userRepo repositories.UserRepository, db *gorm.DB) AuthService {
	return &authService{userRepo: userRepo, db: db}
}

func (s *authService) RegisterUser(user *models.User) (*models.User, error) {
	if existingUser, _ := s.userRepo.FindByEmail(user.Email); existingUser != nil {
		return nil, errors.New("el email ya está registrado")
	}

	if existingUser, _ := s.userRepo.FindByUsername(user.Username); existingUser != nil {
		return nil, errors.New("el nombre de usuario ya está registrado")
	}

	var err error
	user.Password, err = security.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("error al encriptar la contraseña")
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
