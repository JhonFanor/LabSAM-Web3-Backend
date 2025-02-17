package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type DocumentationSubtopicRepository interface {
	Create(documentationSubtopic *models.DocumentationSubtopic) (*models.DocumentationSubtopic, error)
	GetByID(documentationID, subtopicID uint) (*models.DocumentationSubtopic, error)
	Delete(documentationID, subtopicID uint) error
}

type documentationSubtopicRepository struct {
	db *gorm.DB
}

func NewDocumentationSubtopicRepository(db *gorm.DB) DocumentationSubtopicRepository {
	return &documentationSubtopicRepository{
		db: db,
	}
}

func (r *documentationSubtopicRepository) Create(documentationSubtopic *models.DocumentationSubtopic) (*models.DocumentationSubtopic, error) {
	if err := r.db.Create(documentationSubtopic).Error; err != nil {
		return nil, err
	}
	return documentationSubtopic, nil
}

func (r *documentationSubtopicRepository) GetByID(documentationID, subtopicID uint) (*models.DocumentationSubtopic, error) {
	var documentationSubtopic models.DocumentationSubtopic
	if err := r.db.Where("documentation_id = ? AND subtopic_id = ?", documentationID, subtopicID).First(&documentationSubtopic).Error; err != nil {
		return nil, err
	}
	return &documentationSubtopic, nil
}

func (r *documentationSubtopicRepository) Delete(documentationID, subtopicID uint) error {
	return r.db.Where("documentation_id = ? AND subtopic_id = ?", documentationID, subtopicID).
		Delete(&models.DocumentationSubtopic{}).Error
}
