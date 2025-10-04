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
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.Company, error)
	GetRandomApproved() (*models.Company, error)
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
	if err := r.dbManager.Create(company, r.db); err != nil {
		return nil, err
	}
	return company, nil
}

func (r *companyRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.Company{})

	return paginationInfo, nil
}

func (r *companyRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.Company{},
	)

	return paginationInfo, nil
}

func (r *companyRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.Company{},
	)

	return paginationInfo, nil
}

func (r *companyRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Company{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *companyRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("company_subtopic cs").
		Joins("JOIN companies c ON c.id = cs.company_id").
		Where("cs.subtopic_id = ? AND c.is_approved = ?", subtopicID, true).
		Count(&count).Error
	return count, err
}

func (r *companyRepository) GetByID(id uint) (*models.Company, error) {

	var company models.Company
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&company, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Preload("Localitation")); err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) GetRandomApproved() (*models.Company, error) {
	var company models.Company
	err := r.db.
		Preload("User").
		Preload("User.RegularUser").
		Preload("User.UniversityUser").
		Preload("User.BusinessUser").
		Where("is_approved = ?", true).
		Order("RANDOM()").
		First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) Update(company *models.Company, updates map[string]interface{}) error {
	return r.dbManager.Update(company, updates, r.db)
}

func (r *companyRepository) Delete(id uint) error {
	company := models.Company{ID: id}
	return r.dbManager.Delete(&company, r.db)
}
