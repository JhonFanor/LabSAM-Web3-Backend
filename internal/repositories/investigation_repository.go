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
	Update(investigation *models.Investigation) error
	Delete(id uint) error
}

type investigationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewInvestigationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) InvestigationRepository {
	return &investigationRepository{
		db: db,
		qm: qm,
	}
}

func (r *investigationRepository) Create(investigation *models.Investigation) (*models.Investigation, error) {
	if err := r.db.Create(investigation).Error; err != nil {
		return nil, err
	}
	return investigation, nil
}

func (r *investigationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Investigation{})

	return paginationInfo, nil
}

func (r *investigationRepository) GetByID(id uint) (*models.Investigation, error) {
	var investigation models.Investigation
	if err := r.db.First(&investigation, id).Error; err != nil {
		return nil, err
	}
	return &investigation, nil
}

func (r *investigationRepository) Update(investigation *models.Investigation) error {
	return r.db.Save(investigation).Error
}

func (r *investigationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Investigation{}, id).Error
}
