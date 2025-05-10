package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *models.News) (*models.News, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
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
	if err := r.dbManager.Create(news); err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.News{})
	log.Print(paginationInfo.Data)
	return paginationInfo, nil
}

func (r *newsRepository) GetByID(id uint) (*models.News, error) {
	r.dbManager.DB = r.dbManager.DB.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")

	var news models.News
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&news, conditions); err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) Update(news *models.News, updates map[string]interface{}) error {
	return r.dbManager.Update(news, updates)
}

func (r *newsRepository) Delete(id uint) error {
	news := models.News{ID: id}
	return r.dbManager.Delete(&news)
}
