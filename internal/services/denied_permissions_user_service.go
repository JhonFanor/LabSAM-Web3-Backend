package services

import (
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
)

type DeniedPermissionUserService interface {
	GetAllByUser(userID uint) ([]models.DeniedPermissionUser, error)
	Assign(permissionID, userID uint) error
	Revoke(permissionID, userID uint) error
}

type deniedPermissionUserService struct {
	repo repositories.DeniedPermissionUserRepository
}

func NewDeniedPermissionUserService(repo repositories.DeniedPermissionUserRepository) DeniedPermissionUserService {
	return &deniedPermissionUserService{repo: repo}
}

func (s *deniedPermissionUserService) GetAllByUser(userID uint) ([]models.DeniedPermissionUser, error) {
	return s.repo.GetAllByUser(userID)
}

func (s *deniedPermissionUserService) Assign(permissionID, userID uint) error {
	return s.repo.Assign(permissionID, userID)
}

func (s *deniedPermissionUserService) Revoke(permissionID, userID uint) error {
	return s.repo.Revoke(permissionID, userID)
}
