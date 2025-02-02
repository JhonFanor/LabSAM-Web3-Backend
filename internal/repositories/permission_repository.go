package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(permission *models.Permission) (*models.Permission, error)
	Update(permission *models.Permission) error
	Delete(id uint) error
	GetByID(id uint) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
	GetAllByUser(userID uint, roleID uint) ([]models.Permission, error)
}

type permissionRepository struct {
	permissionUserRepository PermissionUserRepository
	permissionRoleRepository PermissionRoleRepository
	db                       *gorm.DB
}

func NewPermissionRepository(permissionUserRepository PermissionUserRepository, permissionRoleRepository PermissionRoleRepository, db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		permissionUserRepository: permissionUserRepository,
		permissionRoleRepository: permissionRoleRepository,
		db:                       db,
	}
}

func (r *permissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	if err := r.db.Create(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *permissionRepository) Update(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Permission{}, id).Error
}

func (r *permissionRepository) GetByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	if err := r.db.First(&permission, id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *permissionRepository) GetAllByUser(userID uint, roleID uint) ([]models.Permission, error) {
	var userPermissions []models.PermissionUser
	var rolePermissions []models.PermissionRole
	var permissions []models.Permission

	userPermissions, err := r.permissionUserRepository.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}

	for _, userPermission := range userPermissions {
		permission, err := r.GetByID(userPermission.PermissionID)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, *permission)
	}

	rolePermissions, err = r.permissionRoleRepository.GetAllByRole(roleID)
	if err != nil {
		return nil, err
	}

	for _, rolePermission := range rolePermissions {
		permission, err := r.GetByID(rolePermission.PermissionID)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, *permission)
	}

	return permissions, nil
}
