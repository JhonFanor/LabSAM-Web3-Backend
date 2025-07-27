package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type RejectionCommentRepository interface {
	Create(comment *models.RejectionComment) (*models.RejectionComment, error)
	Update(comment *models.RejectionComment, updates map[string]interface{}) error
	Delete(id uint) error
	GetAllByResource(resourceType string, resourceID uint) ([]models.RejectionComment, error)
	GetByID(id uint) (*models.RejectionComment, error)
}

type rejectionCommentRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewRejectionCommentRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) RejectionCommentRepository {
	return &rejectionCommentRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *rejectionCommentRepository) Create(comment *models.RejectionComment) (*models.RejectionComment, error) {
	if err := r.dbManager.Create(comment, r.db); err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *rejectionCommentRepository) Update(comment *models.RejectionComment, updates map[string]interface{}) error {
	return r.dbManager.Update(comment, updates, r.db)
}

func (r *rejectionCommentRepository) Delete(id uint) error {
	comment := models.RejectionComment{ID: id}
	return r.dbManager.Delete(&comment, r.db)
}

func (r *rejectionCommentRepository) GetAllByResource(resourceType string, resourceID uint) ([]models.RejectionComment, error) {
	var comments []models.RejectionComment
	err := r.db.
		Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *rejectionCommentRepository) GetByID(id uint) (*models.RejectionComment, error) {
	var rejectionComment models.RejectionComment
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&rejectionComment, conditions, r.db); err != nil {
		return nil, err
	}
	return &rejectionComment, nil
}
