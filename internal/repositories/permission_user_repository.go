package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type PermissionUserRepository interface {
	AssignPermission(permissionUser *models.PermissionUser) error
	RevokePermission(permissionID, userID uint) error
	GetPermission(permissionID, userID uint) (*models.PermissionUser, error)
	GetAllByUser(userID uint) ([]models.PermissionUser, error)
}

type permissionUserRepository struct {
	db *gorm.DB
}

func NewPermissionUserRepository(db *gorm.DB) PermissionUserRepository {
	return &permissionUserRepository{db: db}
}

func (r *permissionUserRepository) AssignPermission(permissionUser *models.PermissionUser) error {
	return r.db.Create(permissionUser).Error
}

func (r *permissionUserRepository) RevokePermission(permissionID, userID uint) error {
	return r.db.Where("permission_id = ? AND user_id = ?", permissionID, userID).Delete(&models.PermissionUser{}).Error
}

func (r *permissionUserRepository) GetPermission(permissionID, userID uint) (*models.PermissionUser, error) {
	var permissionUser models.PermissionUser
	if err := r.db.Where("permission_id = ? AND user_id = ?", permissionID, userID).First(&permissionUser).Error; err != nil {
		return nil, err
	}
	return &permissionUser, nil
}

func (r *permissionUserRepository) GetAllByUser(userID uint) ([]models.PermissionUser, error) {
	var permissions []models.PermissionUser
	if err := r.db.Where("user_id = ?", userID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
