package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type JobExchangeSubtopicRepository interface {
	Create(jobExchangeSubtopic *models.JobExchangeSubtopic) (*models.JobExchangeSubtopic, error)
	Delete(jobExchangeID, subtopicID uint) error
}

type jobExchangeSubtopicRepository struct {
	db *gorm.DB
}

func NewJobExchangeSubtopicRepository(db *gorm.DB) JobExchangeSubtopicRepository {
	return &jobExchangeSubtopicRepository{
		db: db,
	}
}

func (r *jobExchangeSubtopicRepository) Create(jobExchangeSubtopic *models.JobExchangeSubtopic) (*models.JobExchangeSubtopic, error) {
	if err := r.db.Create(jobExchangeSubtopic).Error; err != nil {
		return nil, err
	}
	return jobExchangeSubtopic, nil
}

func (r *jobExchangeSubtopicRepository) Delete(jobExchangeID, subtopicID uint) error {
	return r.db.Where("job_exchange_id = ? AND subtopic_id = ?", jobExchangeID, subtopicID).
		Delete(&models.JobExchangeSubtopic{}).Error
}
