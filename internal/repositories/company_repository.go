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
	Update(company *models.Company) error
	Delete(id uint) error
}

type companyRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewCompanyRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) CompanyRepository {
	return &companyRepository{
		db: db,
		qm: qm,
	}
}

func (r *companyRepository) Create(company *models.Company) (*models.Company, error) {
	if err := r.db.Create(company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (r *companyRepository) GetByID(id uint) (*models.Company, error) {
	var company models.Company
	if err := r.db.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Company{})

	return paginationInfo, nil
}

func (r *companyRepository) Update(company *models.Company) error {
	return r.db.Save(company).Error
}

func (r *companyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Company{}, id).Error
}
