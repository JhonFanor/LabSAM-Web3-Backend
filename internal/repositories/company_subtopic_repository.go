package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type CompanySubtopicRepository interface {
	Create(companySubtopic *models.CompanySubtopic) (*models.CompanySubtopic, error)
	GetByID(companyID, subtopicID uint) (*models.CompanySubtopic, error)
	Delete(companyID, subtopicID uint) error
}

type companySubtopicRepository struct {
	db *gorm.DB
}

func NewCompanySubtopicRepository(db *gorm.DB) CompanySubtopicRepository {
	return &companySubtopicRepository{
		db: db,
	}
}

func (r *companySubtopicRepository) Create(companySubtopic *models.CompanySubtopic) (*models.CompanySubtopic, error) {
	if err := r.db.Create(companySubtopic).Error; err != nil {
		return nil, err
	}
	return companySubtopic, nil
}

func (r *companySubtopicRepository) GetByID(companyID, subtopicID uint) (*models.CompanySubtopic, error) {
	var companySubtopic models.CompanySubtopic
	if err := r.db.Where("company_id = ? AND subtopic_id = ?", companyID, subtopicID).First(&companySubtopic).Error; err != nil {
		return nil, err
	}
	return &companySubtopic, nil
}

func (r *companySubtopicRepository) Delete(companyID, subtopicID uint) error {
	return r.db.Where("company_id = ? AND subtopic_id = ?", companyID, subtopicID).
		Delete(&models.CompanySubtopic{}).Error
}
