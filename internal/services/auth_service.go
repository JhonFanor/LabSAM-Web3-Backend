package services

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type AuthService interface {
	RegisterRegularUser(user *models.User, regularUser *models.RegularUser) (*models.User, error)
	RegisterUniversityUser(user *models.User, universityUser *models.UniversityUser) (*models.User, error)
	RegisterBusinessUser(user *models.User, businessUser *models.BusinessUser) (*models.User, error)
}

type authService struct {
	userService           UserService
	regularUserService    RegularUserService
	universityUserService UniversityUserService
	businessUserService   BusinessUserService
	roleServcie           RoleService
	db                    *gorm.DB
}

func NewAuthService(userService UserService, regularUserService RegularUserService, universityUserService UniversityUserService, businessUserService BusinessUserService, roleServcie RoleService, db *gorm.DB) AuthService {
	return &authService{
		userService:           userService,
		regularUserService:    regularUserService,
		universityUserService: universityUserService,
		businessUserService:   businessUserService,
		roleServcie:           roleServcie,
		db:                    db,
	}
}

func (s *authService) RegisterRegularUser(user *models.User, regularUser *models.RegularUser) (*models.User, error) {
	role, err := s.roleServcie.GetByName("business")
	if err != nil {
		return nil, err
	}

	user.RoleID = role.ID

	createdUser, err := s.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	regularUser.UserID = createdUser.ID
	_, err = s.regularUserService.CreateRegularUser(regularUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *authService) RegisterUniversityUser(user *models.User, universityUser *models.UniversityUser) (*models.User, error) {
	role, err := s.roleServcie.GetByName("business")
	if err != nil {
		return nil, err
	}

	user.RoleID = role.ID

	createdUser, err := s.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	universityUser.UserID = createdUser.ID
	_, err = s.universityUserService.CreateUniversityUser(universityUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *authService) RegisterBusinessUser(user *models.User, businessUser *models.BusinessUser) (*models.User, error) {

	role, err := s.roleServcie.GetByName("business")
	if err != nil {
		return nil, err
	}

	user.RoleID = role.ID

	createdUser, err := s.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	businessUser.UserID = createdUser.ID
	_, err = s.businessUserService.CreateBusinessUser(businessUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
