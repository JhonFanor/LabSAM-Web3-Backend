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
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.Investigation, error)
	GetRandomApproved() (*models.Investigation, error)
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
	if err := r.dbManager.Create(investigation, r.db); err != nil {
		return nil, err
	}
	return investigation, nil
}

func (r *investigationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.Investigation{})

	return paginationInfo, nil
}

func (r *investigationRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.Investigation{},
	), nil
}

func (r *investigationRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.Investigation{},
	), nil
}

func (r *investigationRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Investigation{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *investigationRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("investigation_subtopic its").
		Joins("JOIN investigations i ON i.id = its.investigation_id").
		Where("its.subtopic_id = ? AND i.is_approved = ?", subtopicID, true).
		Count(&count).Error
	return count, err
}

func (r *investigationRepository) GetByID(id uint) (*models.Investigation, error) {
	var investigation models.Investigation
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&investigation, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &investigation, nil
}

func (r *investigationRepository) GetRandomApproved() (*models.Investigation, error) {
	var investigation models.Investigation
	err := r.db.
		Preload("User").
		Preload("User.RegularUser").
		Preload("User.UniversityUser").
		Preload("User.BusinessUser").
		Where("is_approved = ?", true).
		Order("RANDOM()").
		First(&investigation).Error
	if err != nil {
		return nil, err
	}
	return &investigation, nil
}

func (r *investigationRepository) Update(investigation *models.Investigation, updates map[string]interface{}) error {
	return r.dbManager.Update(investigation, updates, r.db)
}

func (r *investigationRepository) Delete(id uint) error {
	investigation := models.Investigation{ID: id}
	return r.dbManager.Delete(&investigation, r.db)
}
