package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type NewsSubtopicRepository interface {
	Create(newsSubtopic *models.NewsSubtopic) (*models.NewsSubtopic, error)
	GetByID(newsID, subtopicID uint) (*models.NewsSubtopic, error)
	Delete(newsID, subtopicID uint) error
}

type newsSubtopicRepository struct {
	db *gorm.DB
}

func NewNewsSubtopicRepository(db *gorm.DB) NewsSubtopicRepository {
	return &newsSubtopicRepository{
		db: db,
	}
}

func (r *newsSubtopicRepository) Create(newsSubtopic *models.NewsSubtopic) (*models.NewsSubtopic, error) {
	if err := r.db.Create(newsSubtopic).Error; err != nil {
		return nil, err
	}
	return newsSubtopic, nil
}

func (r *newsSubtopicRepository) GetByID(newsID, subtopicID uint) (*models.NewsSubtopic, error) {
	var newsSubtopic models.NewsSubtopic
	if err := r.db.Where("news_id = ? AND subtopic_id = ?", newsID, subtopicID).First(&newsSubtopic).Error; err != nil {
		return nil, err
	}
	return &newsSubtopic, nil
}

func (r *newsSubtopicRepository) Delete(newsID, subtopicID uint) error {
	return r.db.Where("new_id = ? AND subtopic_id = ?", newsID, subtopicID).
		Delete(&models.NewsSubtopic{}).Error
}
