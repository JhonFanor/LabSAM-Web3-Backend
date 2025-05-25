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
	GetByID(id uint) (*models.BankOfResume, error)
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

func (r *bankOfResumeRepository) GetByID(id uint) (*models.BankOfResume, error) {
	var bankOfResume models.BankOfResume
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&bankOfResume, conditions, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
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
