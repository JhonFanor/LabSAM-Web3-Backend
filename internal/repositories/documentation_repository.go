package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentationRepository interface {
	Create(documentation *models.Documentation) (*models.Documentation, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	GetByID(id uint) (*models.Documentation, error)
	Update(documentation *models.Documentation, updates map[string]interface{}) error
	Delete(id uint) error
}

type documentationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewDocumentationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) DocumentationRepository {
	return &documentationRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *documentationRepository) Create(documentation *models.Documentation) (*models.Documentation, error) {
	if err := r.dbManager.Create(documentation, r.db); err != nil {
		return nil, err
	}
	return documentation, nil
}

func (r *documentationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.Documentation{})

	return paginationInfo, nil
}

func (r *documentationRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.Documentation{},
	)

	return paginationInfo, nil
}

func (r *documentationRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.Documentation{},
	)

	return paginationInfo, nil
}

func (r *documentationRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Documentation{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *documentationRepository) GetByID(id uint) (*models.Documentation, error) {

	var documentation models.Documentation
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&documentation, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true)); err != nil {
		return nil, err
	}
	return &documentation, nil
}

func (r *documentationRepository) Update(documentation *models.Documentation, updates map[string]interface{}) error {
	return r.dbManager.Update(documentation, updates, r.db)
}

func (r *documentationRepository) Delete(id uint) error {
	documentation := models.Documentation{ID: id}
	return r.dbManager.Delete(&documentation, r.db)
}
