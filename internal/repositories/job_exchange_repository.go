package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type JobExchangeRepository interface {
	Create(jobExchange *models.JobExchange) (*models.JobExchange, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.JobExchange, error)
	Update(jobExchange *models.JobExchange, updates map[string]interface{}) error
	Delete(id uint) error
}

type jobExchangeRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
}

func NewJobExchangeRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) JobExchangeRepository {
	return &jobExchangeRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *jobExchangeRepository) Create(jobExchange *models.JobExchange) (*models.JobExchange, error) {
	if err := r.dbManager.Create(jobExchange); err != nil {
		return nil, err
	}
	return jobExchange, nil
}

func (r *jobExchangeRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	r.qm.DB = r.qm.DB.Preload("User")
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.JobExchange{})

	return paginationInfo, nil
}

func (r *jobExchangeRepository) GetByID(id uint) (*models.JobExchange, error) {
	r.dbManager.DB = r.qm.DB.Preload("Subtopics").Preload("User")

	var jobExchange models.JobExchange
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&jobExchange, conditions); err != nil {
		return nil, err
	}
	return &jobExchange, nil
}

func (r *jobExchangeRepository) Update(jobExchange *models.JobExchange, updates map[string]interface{}) error {
	return r.dbManager.Update(jobExchange, updates)
}

func (r *jobExchangeRepository) Delete(id uint) error {
	jobExchange := models.JobExchange{ID: id}
	return r.dbManager.Delete(&jobExchange)
}
