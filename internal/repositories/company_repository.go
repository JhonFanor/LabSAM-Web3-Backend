package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company *models.Company) (*models.Company, error)
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
