package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type SubtopicRepository interface {
	Create(subtopic *models.Subtopic) (*models.Subtopic, error)
	Update(subtopic *models.Subtopic) error
	Delete(id uint) error
	GetByID(id uint) (*models.Subtopic, error)
	GetAll() ([]models.Subtopic, error)
}

type subtopicRepository struct {
	db *gorm.DB
}

func NewSubtopicRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) SubtopicRepository {
	return &subtopicRepository{db: db}
}

func (r *subtopicRepository) Create(subtopic *models.Subtopic) (*models.Subtopic, error) {
	if err := r.db.Create(subtopic).Error; err != nil {
		return nil, err
	}
	return subtopic, nil
}

func (r *subtopicRepository) Update(subtopic *models.Subtopic) error {
	return r.db.Save(subtopic).Error
}

func (r *subtopicRepository) Delete(id uint) error {
	return r.db.Delete(&models.Subtopic{}, id).Error
}

func (r *subtopicRepository) GetByID(id uint) (*models.Subtopic, error) {
	var subtopic models.Subtopic
	if err := r.db.First(&subtopic, id).Error; err != nil {
		return nil, err
	}
	return &subtopic, nil
}

func (r *subtopicRepository) GetAll() ([]models.Subtopic, error) {
	var subtopics []models.Subtopic

	if err := r.db.Find(&subtopics).Error; err != nil {
		return nil, err
	}

	return subtopics, nil
}
