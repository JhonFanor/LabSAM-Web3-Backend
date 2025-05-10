package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InvestigationRepository interface {
	Create(investigation *models.Investigation) (*models.Investigation, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.Investigation, error)
	Update(investigation *models.Investigation, updates map[string]interface{}) error
	Delete(id uint) error
}

type investigationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewInvestigationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) InvestigationRepository {
	return &investigationRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *investigationRepository) Create(investigation *models.Investigation) (*models.Investigation, error) {
	if err := r.dbManager.Create(investigation); err != nil {
		return nil, err
	}
	return investigation, nil
}

func (r *investigationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Where("is_approved = ?", true), &models.Investigation{})

	return paginationInfo, nil
}

func (r *investigationRepository) GetByID(id uint) (*models.Investigation, error) {
	r.dbManager.DB = r.dbManager.DB.Preload("Subtopics").Preload("User")

	var investigation models.Investigation
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&investigation, conditions); err != nil {
		return nil, err
	}
	return &investigation, nil
}

func (r *investigationRepository) Update(investigation *models.Investigation, updates map[string]interface{}) error {
	return r.dbManager.Update(investigation, updates)
}

func (r *investigationRepository) Delete(id uint) error {
	investigation := models.Investigation{ID: id}
	return r.dbManager.Delete(&investigation)
}
