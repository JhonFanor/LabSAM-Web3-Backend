package services

import (
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(c *gin.Context) (*dto.PaginationDTO, error)
	FindUserByEmail(email string) (*models.User, error)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *models.User) error {

	var err error

	updates := utils.StructToMap(user)

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	}

	if len(updates) == 0 {
		return nil
	}

	existing, err := s.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	existing.BusinessUser = nil
	existing.RegularUser = nil
	existing.UniversityUser = nil
	return s.repo.Update(existing, updates)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAllUsers(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *userService) FindUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
