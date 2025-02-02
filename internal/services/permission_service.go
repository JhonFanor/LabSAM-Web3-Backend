package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type PermissionService interface {
	CreatePermission(permission *models.Permission) (*models.Permission, error)
	UpdatePermission(permission *models.Permission) error
	DeletePermission(id uint) error
	GetPermissionByID(id uint) (*models.Permission, error)
	GetAllPermissions() ([]models.Permission, error)
	GetAllPermissionsByUser(userID uint, roleID uint) ([]models.Permission, error)
}

type permissionService struct {
	repo repositories.PermissionRepository
	db   *gorm.DB
}

func NewPermissionService(repo repositories.PermissionRepository, db *gorm.DB) PermissionService {
	return &permissionService{
		repo: repo,
		db:   db,
	}
}

func (s *permissionService) CreatePermission(permission *models.Permission) (*models.Permission, error) {
	return s.repo.Create(permission)
}

func (s *permissionService) UpdatePermission(permission *models.Permission) error {
	return s.repo.Update(permission)
}

func (s *permissionService) DeletePermission(id uint) error {
	return s.repo.Delete(id)
}

func (s *permissionService) GetPermissionByID(id uint) (*models.Permission, error) {
	return s.repo.GetByID(id)
}

func (s *permissionService) GetAllPermissions() ([]models.Permission, error) {
	return s.repo.GetAll()
}

func (s *permissionService) GetAllPermissionsByUser(userID uint, roleID uint) ([]models.Permission, error) {
	return s.repo.GetAllByUser(userID, roleID)
}
