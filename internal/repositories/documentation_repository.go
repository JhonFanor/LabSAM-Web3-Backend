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
	GetByID(id uint) (*models.Documentation, error)
	Update(documentation *models.Documentation) error
	Delete(id uint) error
}

type documentationRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewDocumentationRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) DocumentationRepository {
	return &documentationRepository{
		db: db,
		qm: qm,
	}
}

func (r *documentationRepository) Create(documentation *models.Documentation) (*models.Documentation, error) {
	if err := r.db.Create(documentation).Error; err != nil {
		return nil, err
	}
	return documentation, nil
}

func (r *documentationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Documentation{})

	return paginationInfo, nil
}

func (r *documentationRepository) GetByID(id uint) (*models.Documentation, error) {
	var documentation models.Documentation
	if err := r.db.First(&documentation, id).Error; err != nil {
		return nil, err
	}
	return &documentation, nil
}

func (r *documentationRepository) Update(documentation *models.Documentation) error {
	return r.db.Save(documentation).Error
}

func (r *documentationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Documentation{}, id).Error
}
