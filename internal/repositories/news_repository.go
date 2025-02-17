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
	GetByID(id uint) (*models.News, error)
	Update(news *models.News) error
	Delete(id uint) error
}

type newsRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewNewsRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) NewsRepository {
	return &newsRepository{
		db: db,
		qm: qm,
	}
}

func (r *newsRepository) Create(news *models.News) (*models.News, error) {
	if err := r.db.Create(news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.News{})

	return paginationInfo, nil
}

func (r *newsRepository) GetByID(id uint) (*models.News, error) {
	var news models.News
	if err := r.db.First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) Update(news *models.News) error {
	return r.db.Save(news).Error
}

func (r *newsRepository) Delete(id uint) error {
	return r.db.Delete(&models.News{}, id).Error
}
