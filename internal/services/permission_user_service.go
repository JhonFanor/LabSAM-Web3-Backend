package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type PermissionUserService interface {
	AssignPermissionToUser(permissionUser *models.PermissionUser) error
	RevokePermissionFromUser(permissionID, userID uint) error
	GetPermissionUser(permissionID, userID uint) (*models.PermissionUser, error)
	GetAllPermissionsByUser(userID uint) ([]models.PermissionUser, error)
}

type permissionUserService struct {
	repo repositories.PermissionUserRepository
	db   *gorm.DB
}

func NewPermissionUserService(repo repositories.PermissionUserRepository, db *gorm.DB) PermissionUserService {
	return &permissionUserService{
		repo: repo,
		db:   db,
	}
}

func (s *permissionUserService) AssignPermissionToUser(permissionUser *models.PermissionUser) error {
	return s.repo.AssignPermission(permissionUser)
}

func (s *permissionUserService) RevokePermissionFromUser(permissionID, userID uint) error {
	return s.repo.RevokePermission(permissionID, userID)
}

func (s *permissionUserService) GetPermissionUser(permissionID, userID uint) (*models.PermissionUser, error) {
	return s.repo.GetPermission(permissionID, userID)
}

func (s *permissionUserService) GetAllPermissionsByUser(userID uint) ([]models.PermissionUser, error) {
	return s.repo.GetAllByUser(userID)
}
