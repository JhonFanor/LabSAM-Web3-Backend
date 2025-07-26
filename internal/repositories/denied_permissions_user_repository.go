package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type DeniedPermissionUserRepository interface {
	GetAllByUser(userID uint) ([]models.DeniedPermissionUser, error)
	Assign(permissionID, userID uint) error
	Revoke(permissionID, userID uint) error
}

type deniedPermissionUserRepository struct {
	db *gorm.DB
}

func NewDeniedPermissionUserRepository(db *gorm.DB) DeniedPermissionUserRepository {
	return &deniedPermissionUserRepository{db: db}
}

func (r *deniedPermissionUserRepository) GetAllByUser(userID uint) ([]models.DeniedPermissionUser, error) {
	var denied []models.DeniedPermissionUser
	if err := r.db.Where("user_id = ?", userID).Find(&denied).Error; err != nil {
		return nil, err
	}
	return denied, nil
}

func (r *deniedPermissionUserRepository) Assign(permissionID, userID uint) error {
	denied := models.DeniedPermissionUser{
		PermissionID: permissionID,
		UserID:       userID,
	}
	return r.db.Create(&denied).Error
}

func (r *deniedPermissionUserRepository) Revoke(permissionID, userID uint) error {
	return r.db.Delete(&models.DeniedPermissionUser{}, "permission_id = ? AND user_id = ?", permissionID, userID).Error
}
