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
	GetAllByRoleExcludingDenied(roleID, userID uint) ([]models.Permission, error)
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

func (r *permissionRoleRepository) GetAllByRoleExcludingDenied(roleID, userID uint) ([]models.Permission, error) {
	var permissions []models.Permission

	err := r.db.
		Table("permissions").
		Select("permissions.*").
		Joins("JOIN permission_rol ON permission_rol.permission_id = permissions.id").
		Joins("LEFT JOIN denied_permissions_user ON denied_permissions_user.permission_id = permissions.id AND denied_permissions_user.user_id = ?", userID).
		Where("permission_rol.role_id = ?", roleID).
		Where("denied_permissions_user.user_id IS NULL").
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
