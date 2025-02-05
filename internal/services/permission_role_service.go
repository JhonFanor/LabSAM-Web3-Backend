package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type PermissionRoleService interface {
	AssignPermissionToRole(permissionRole *models.PermissionRole) error
	RevokePermissionFromRole(permissionID, roleID uint) error
	GetPermissionRole(permissionID, roleID uint) (*models.PermissionRole, error)
	GetAllPermissionsByRole(roleID uint) ([]models.PermissionRole, error)
}

type permissionRoleService struct {
	repo repositories.PermissionRoleRepository
	db   *gorm.DB
}

func NewPermissionRoleService(repo repositories.PermissionRoleRepository, db *gorm.DB) PermissionRoleService {
	return &permissionRoleService{
		repo: repo,
		db:   db,
	}
}

func (s *permissionRoleService) AssignPermissionToRole(permissionRole *models.PermissionRole) error {
	return s.repo.AssignPermission(permissionRole)
}

func (s *permissionRoleService) RevokePermissionFromRole(permissionID, roleID uint) error {
	return s.repo.RevokePermission(permissionID, roleID)
}

func (s *permissionRoleService) GetPermissionRole(permissionID, roleID uint) (*models.PermissionRole, error) {
	return s.repo.GetPermission(permissionID, roleID)
}

func (s *permissionRoleService) GetAllPermissionsByRole(roleID uint) ([]models.PermissionRole, error) {
	return s.repo.GetAllByRole(roleID)
}
