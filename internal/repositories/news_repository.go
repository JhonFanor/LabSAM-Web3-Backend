package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *models.News) (*models.News, error)
	Delete(newsID, subtopicID uint) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

func (r *newsRepository) Create(news *models.News) (*models.News, error) {
	if err := r.db.Create(news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) Delete(newsID, subtopicID uint) error {
	return r.db.Where("news_id = ? AND subtopic_id = ?", newsID, subtopicID).
		Delete(&models.News{}).Error
}
