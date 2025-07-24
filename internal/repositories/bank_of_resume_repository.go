package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankOfResumeRepository interface {
	Create(bankOfResume *models.BankOfResume) (*models.BankOfResume, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.BankOfResume, error)
	GetRandomApproved() (*models.BankOfResume, error)
	Update(bankOfResume *models.BankOfResume, updates map[string]interface{}) error
	Delete(id uint) error
}

type bankOfResumeRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewBankOfResumeRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) BankOfResumeRepository {
	return &bankOfResumeRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *bankOfResumeRepository) Create(bankOfResume *models.BankOfResume) (*models.BankOfResume, error) {
	if err := r.dbManager.Create(bankOfResume, r.db); err != nil {
		return nil, err
	}
	return bankOfResume, nil
}

func (r *bankOfResumeRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.BankOfResume{})

	return paginationInfo, nil
}

func (r *bankOfResumeRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.BankOfResume{},
	)

	return paginationInfo, nil
}

func (r *bankOfResumeRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.BankOfResume{},
	)

	return paginationInfo, nil
}

func (r *bankOfResumeRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.BankOfResume{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *bankOfResumeRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("bank_of_resume_subtopic").
		Where("subtopic_id = ?", subtopicID).
		Count(&count).Error
	return count, err
}

func (r *bankOfResumeRepository) GetByID(id uint) (*models.BankOfResume, error) {
	var bankOfResume models.BankOfResume
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&bankOfResume, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &bankOfResume, nil
}

func (r *bankOfResumeRepository) GetRandomApproved() (*models.BankOfResume, error) {
	var bankOfResume models.BankOfResume
	err := r.db.
		Preload("User").
		Where("is_approved = ?", true).
		Order("RANDOM()").
		First(&bankOfResume).Error
	if err != nil {
		return nil, err
	}
	return &bankOfResume, nil
}

func (r *bankOfResumeRepository) Update(bankOfResume *models.BankOfResume, updates map[string]interface{}) error {
	return r.dbManager.Update(bankOfResume, updates, r.db)
}

func (r *bankOfResumeRepository) Delete(id uint) error {
	bankOfResume := models.BankOfResume{ID: id}
	return r.dbManager.Delete(&bankOfResume, r.db)
}
