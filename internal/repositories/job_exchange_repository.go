package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type JobExchangeRepository interface {
	Create(jobExchange *models.JobExchange) (*models.JobExchange, error)
}

type jobExchangeRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewJobExchangeRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) JobExchangeRepository {
	return &jobExchangeRepository{
		db: db,
		qm: qm,
	}
}

func (r *jobExchangeRepository) Create(jobExchange *models.JobExchange) (*models.JobExchange, error) {
	if err := r.db.Create(jobExchange).Error; err != nil {
		return nil, err
	}
	return jobExchange, nil
}
