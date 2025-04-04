package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type JobBoardSubtopicRepository interface {
	Create(jobBoardSubtopic *models.JobBoardSubtopic) (*models.JobBoardSubtopic, error)
	GetByID(jobBoardID, subtopicID uint) (*models.JobBoardSubtopic, error)
	Delete(jobBoardID, subtopicID uint) error
}

type jobBoardSubtopicRepository struct {
	db *gorm.DB
}

func NewJobBoardSubtopicRepository(db *gorm.DB) JobBoardSubtopicRepository {
	return &jobBoardSubtopicRepository{
		db: db,
	}
}

func (r *jobBoardSubtopicRepository) Create(jobBoardSubtopic *models.JobBoardSubtopic) (*models.JobBoardSubtopic, error) {
	if err := r.db.Create(jobBoardSubtopic).Error; err != nil {
		return nil, err
	}
	return jobBoardSubtopic, nil
}

func (r *jobBoardSubtopicRepository) GetByID(jobBoardID, subtopicID uint) (*models.JobBoardSubtopic, error) {
	var jobBoardSubtopic models.JobBoardSubtopic
	if err := r.db.Where("job_exchange_id = ? AND subtopic_id = ?", jobBoardID, subtopicID).First(&jobBoardSubtopic).Error; err != nil {
		return nil, err
	}
	return &jobBoardSubtopic, nil
}

func (r *jobBoardSubtopicRepository) Delete(jobBoardID, subtopicID uint) error {
	return r.db.Where("job_exchange_id = ? AND subtopic_id = ?", jobBoardID, subtopicID).
		Delete(&models.JobBoardSubtopic{}).Error
}
