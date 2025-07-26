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
	GetAllAssignableToUser(userID uint, roleID uint) ([]models.Permission, error)
}

type permissionRepository struct {
	permissionUserRepository       PermissionUserRepository
	permissionRoleRepository       PermissionRoleRepository
	deniedPermissionUserRepository DeniedPermissionUserRepository
	db                             *gorm.DB
}

func NewPermissionRepository(
	permissionUserRepo PermissionUserRepository,
	permissionRoleRepo PermissionRoleRepository,
	deniedPermissionRepo DeniedPermissionUserRepository,
	db *gorm.DB,
) PermissionRepository {
	return &permissionRepository{
		permissionUserRepository:       permissionUserRepo,
		permissionRoleRepository:       permissionRoleRepo,
		deniedPermissionUserRepository: deniedPermissionRepo,
		db:                             db,
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
	var permissions []models.Permission
	permMap := make(map[uint]models.Permission)

	userPerms, err := r.permissionUserRepository.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}
	for _, up := range userPerms {
		p, err := r.GetByID(up.PermissionID)
		if err != nil {
			return nil, err
		}
		permMap[p.ID] = *p
	}

	rolePerms, err := r.permissionRoleRepository.GetAllByRole(roleID)
	if err != nil {
		return nil, err
	}
	for _, rp := range rolePerms {
		p, err := r.GetByID(rp.PermissionID)
		if err != nil {
			return nil, err
		}
		permMap[p.ID] = *p
	}

	deniedPerms, err := r.deniedPermissionUserRepository.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}
	for _, dp := range deniedPerms {
		delete(permMap, dp.PermissionID)
	}

	// Convertir map a slice
	for _, p := range permMap {
		permissions = append(permissions, p)
	}

	return permissions, nil
}

func (r *permissionRepository) GetAllAssignableToUser(userID uint, roleID uint) ([]models.Permission, error) {
	allPerms, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	userPerms, err := r.permissionUserRepository.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}
	userPermSet := make(map[uint]struct{})
	for _, up := range userPerms {
		userPermSet[up.PermissionID] = struct{}{}
	}

	rolePerms, err := r.permissionRoleRepository.GetAllByRole(roleID)
	if err != nil {
		return nil, err
	}
	rolePermSet := make(map[uint]struct{})
	for _, rp := range rolePerms {
		rolePermSet[rp.PermissionID] = struct{}{}
	}

	deniedPerms, err := r.deniedPermissionUserRepository.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}
	deniedPermSet := make(map[uint]struct{})
	for _, dp := range deniedPerms {
		deniedPermSet[dp.PermissionID] = struct{}{}
	}

	var assignable []models.Permission
	for _, p := range allPerms {
		_, inUser := userPermSet[p.ID]
		_, inRole := rolePermSet[p.ID]
		_, denied := deniedPermSet[p.ID]

		if !inUser && !inRole && !denied {
			assignable = append(assignable, p)
		}
	}

	return assignable, nil
}
