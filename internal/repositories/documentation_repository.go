package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type DocumentationRepository interface {
	Create(documentation *models.Documentation) (*models.Documentation, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.Documentation, error)
	Update(documentation *models.Documentation, updates map[string]interface{}) error
	Delete(id uint) error
}

type documentationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
}

func NewDocumentationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) DocumentationRepository {
	return &documentationRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *documentationRepository) Create(documentation *models.Documentation) (*models.Documentation, error) {
	if err := r.dbManager.Create(documentation); err != nil {
		return nil, err
	}
	return documentation, nil
}

func (r *documentationRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	r.qm.DB = r.qm.DB.Preload("User")
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Documentation{})

	return paginationInfo, nil
}

func (r *documentationRepository) GetByID(id uint) (*models.Documentation, error) {
	r.dbManager.DB = r.qm.DB.Preload("Subtopics").Preload("User")

	var documentation models.Documentation
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&documentation, conditions); err != nil {
		return nil, err
	}
	return &documentation, nil
}

func (r *documentationRepository) Update(documentation *models.Documentation, updates map[string]interface{}) error {
	return r.dbManager.Update(documentation, updates)
}

func (r *documentationRepository) Delete(id uint) error {
	documentation := models.Documentation{ID: id}
	return r.dbManager.Delete(&documentation)
}
