package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LegislationRepository interface {
	Create(legislation *models.Legislation) (*models.Legislation, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.Legislation, error)
	Update(legislation *models.Legislation) error
	Delete(id uint) error
}

type legislationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewLegislationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) LegislationRepository {
	return &legislationRepository{
		db: db,
		qm: qm,
	}
}

func (r *legislationRepository) Create(legislation *models.Legislation) (*models.Legislation, error) {
	if err := r.db.Create(legislation).Error; err != nil {
		return nil, err
	}
	return legislation, nil
}

func (r *legislationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Legislation{})

	return paginationInfo, nil
}

func (r *legislationRepository) GetByID(id uint) (*models.Legislation, error) {
	var legislation models.Legislation
	if err := r.db.First(&legislation, id).Error; err != nil {
		return nil, err
	}
	return &legislation, nil
}

func (r *legislationRepository) Update(legislation *models.Legislation) error {
	return r.db.Save(legislation).Error
}

func (r *legislationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Legislation{}, id).Error
}
