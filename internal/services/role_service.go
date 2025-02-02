package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type RoleService interface {
	CreateRole(role *models.Role) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
	GetRoleByID(id uint) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
}

type roleService struct {
	repo repositories.RoleRepository
	db   *gorm.DB
}

func NewRoleService(repo repositories.RoleRepository, db *gorm.DB) RoleService {
	return &roleService{
		repo: repo,
		db:   db,
	}
}

func (s *roleService) CreateRole(role *models.Role) (*models.Role, error) {
	return s.repo.Create(role)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.repo.Update(role)
}

func (s *roleService) DeleteRole(id uint) error {
	return s.repo.Delete(id)
}

func (s *roleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) GetByName(name string) (*models.Role, error) {
	return s.repo.GetByName(name)
}

func (s *roleService) GetAllRoles() ([]models.Role, error) {
	return s.repo.GetAll()
}
