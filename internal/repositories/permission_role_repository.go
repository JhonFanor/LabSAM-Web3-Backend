package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type PermissionRoleRepository interface {
	AssignPermission(permissionRole *models.PermissionRole) error
	RevokePermission(permissionID, roleID uint) error
	GetPermission(permissionID, roleID uint) (*models.PermissionRole, error)
	GetAllByRole(roleID uint) ([]models.PermissionRole, error)
}

type permissionRoleRepository struct {
	db *gorm.DB
}

func NewPermissionRoleRepository(db *gorm.DB) PermissionRoleRepository {
	return &permissionRoleRepository{db: db}
}

func (r *permissionRoleRepository) AssignPermission(permissionRole *models.PermissionRole) error {
	return r.db.Create(permissionRole).Error
}

func (r *permissionRoleRepository) RevokePermission(permissionID, roleID uint) error {
	return r.db.Where("permission_id = ? AND role_id = ?", permissionID, roleID).Delete(&models.PermissionRole{}).Error
}

func (r *permissionRoleRepository) GetPermission(permissionID, roleID uint) (*models.PermissionRole, error) {
	var permissionRole models.PermissionRole
	if err := r.db.Where("permission_id = ? AND role_id = ?", permissionID, roleID).First(&permissionRole).Error; err != nil {
		return nil, err
	}
	return &permissionRole, nil
}

func (r *permissionRoleRepository) GetAllByRole(roleID uint) ([]models.PermissionRole, error) {
	var permissions []models.PermissionRole
	if err := r.db.Where("role_id = ?", roleID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
