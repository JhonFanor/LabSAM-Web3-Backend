package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *models.News) (*models.News, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.News, error)
	Update(news *models.News, updates map[string]interface{}) error
	Delete(id uint) error
}

type newsRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewNewsRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) NewsRepository {
	return &newsRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *newsRepository) Create(news *models.News) (*models.News, error) {
	if err := r.dbManager.Create(news, r.db); err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.News{})
	return paginationInfo, nil
}

func (r *newsRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.News{},
	), nil
}

func (r *newsRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.News{},
	), nil
}

func (r *newsRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.News{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *newsRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("news_subtopic").
		Where("subtopic_id = ?", subtopicID).
		Count(&count).Error
	return count, err
}

func (r *newsRepository) GetByID(id uint) (*models.News, error) {
	var news models.News
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&news, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) Update(news *models.News, updates map[string]interface{}) error {
	return r.dbManager.Update(news, updates, r.db)
}

func (r *newsRepository) Delete(id uint) error {
	news := models.News{ID: id}
	return r.dbManager.Delete(&news, r.db)
}
