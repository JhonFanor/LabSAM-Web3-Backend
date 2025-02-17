package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobExchangeRepository interface {
	Create(jobExchange *models.JobExchange) (*models.JobExchange, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.JobExchange, error)
	Update(jobExchange *models.JobExchange) error
	Delete(id uint) error
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

func (r *jobExchangeRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.JobExchange{})

	return paginationInfo, nil
}

func (r *jobExchangeRepository) GetByID(id uint) (*models.JobExchange, error) {
	var jobExchange models.JobExchange
	if err := r.db.First(&jobExchange, id).Error; err != nil {
		return nil, err
	}
	return &jobExchange, nil
}

func (r *jobExchangeRepository) Update(jobExchange *models.JobExchange) error {
	return r.db.Save(jobExchange).Error
}

func (r *jobExchangeRepository) Delete(id uint) error {
	return r.db.Delete(&models.JobExchange{}, id).Error
}
