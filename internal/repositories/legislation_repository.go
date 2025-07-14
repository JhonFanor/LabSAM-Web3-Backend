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
	Update(legislation *models.Legislation, updates map[string]interface{}) error
	Delete(id uint) error
}

type legislationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewLegislationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) LegislationRepository {
	return &legislationRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *legislationRepository) Create(legislation *models.Legislation) (*models.Legislation, error) {
	if err := r.dbManager.Create(legislation, r.db); err != nil {
		return nil, err
	}
	return legislation, nil
}

func (r *legislationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.Legislation{})

	return paginationInfo, nil
}

func (r *legislationRepository) GetByID(id uint) (*models.Legislation, error) {
	var legislation models.Legislation
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&legislation, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &legislation, nil
}

func (r *legislationRepository) Update(legislation *models.Legislation, updates map[string]interface{}) error {
	return r.dbManager.Update(legislation, updates, r.db)
}

func (r *legislationRepository) Delete(id uint) error {
	legislation := models.Legislation{ID: id}
	return r.dbManager.Delete(&legislation, r.db)
}
