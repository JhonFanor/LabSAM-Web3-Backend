package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"
)

type RejectionCommentService interface {
	CreateRejectionComment(comment *models.RejectionComment) (*models.RejectionComment, error)
	UpdateRejectionComment(rejectionComment *models.RejectionComment) error
	DeleteRejectionComment(commentID uint) error
	GetAllByResource(resourceType string, resourceID uint) ([]models.RejectionComment, error)
	GetRejectionCommetByID(id uint) (*models.RejectionComment, error)
}

type rejectionCommentService struct {
	repo repositories.RejectionCommentRepository
}

func NewRejectionCommentService(repo repositories.RejectionCommentRepository) RejectionCommentService {
	return &rejectionCommentService{
		repo: repo,
	}
}

func (s *rejectionCommentService) CreateRejectionComment(comment *models.RejectionComment) (*models.RejectionComment, error) {
	if comment == nil {
		return nil, customerrors.ErrInvalidData
	}

	return s.repo.Create(comment)
}

func (s *rejectionCommentService) UpdateRejectionComment(rejectionComment *models.RejectionComment) error {
	if rejectionComment == nil || rejectionComment.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetRejectionCommetByID(rejectionComment.ID)
	if err != nil {
		return err
	}

	updates := utils.StructToMap(rejectionComment)
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(existing, updates)
}

func (s *rejectionCommentService) DeleteRejectionComment(commentID uint) error {
	if commentID == 0 {
		return customerrors.ErrInvalidID
	}

	return s.repo.Delete(commentID)
}

func (s *rejectionCommentService) GetAllByResource(resourceType string, resourceID uint) ([]models.RejectionComment, error) {
	if resourceType == "" || resourceID == 0 {
		return nil, customerrors.ErrInvalidData
	}

	return s.repo.GetAllByResource(resourceType, resourceID)
}

func (s *rejectionCommentService) GetRejectionCommetByID(id uint) (*models.RejectionComment, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	rejectionComment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return rejectionComment, nil
}
