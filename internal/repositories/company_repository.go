package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company *models.Company) (*models.Company, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.Company, error)
	Update(company *models.Company, updates map[string]interface{}) error
	Delete(id uint) error
}

type companyRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewCompanyRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) CompanyRepository {
	return &companyRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *companyRepository) Create(company *models.Company) (*models.Company, error) {
	if err := r.dbManager.Create(company); err != nil {
		return nil, err
	}
	return company, nil
}

func (r *companyRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Where("is_approved = ?", true), &models.Company{})

	return paginationInfo, nil
}

func (r *companyRepository) GetByID(id uint) (*models.Company, error) {
	r.dbManager.DB = r.dbManager.DB.Preload("Subtopics").Preload("User")

	var company models.Company
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&company, conditions); err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) Update(company *models.Company, updates map[string]interface{}) error {
	return r.dbManager.Update(company, updates)
}

func (r *companyRepository) Delete(id uint) error {
	company := models.Company{ID: id}
	return r.dbManager.Delete(&company)
}
