package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
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
}

func NewBankOfResumeRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) BankOfResumeRepository {
	return &bankOfResumeRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *bankOfResumeRepository) Create(bankOfResume *models.BankOfResume) (*models.BankOfResume, error) {
	if err := r.dbManager.Create(bankOfResume); err != nil {
		return nil, err
	}
	return bankOfResume, nil
}

func (r *bankOfResumeRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	r.qm.DB = r.qm.DB.Preload("User")
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.BankOfResume{})

	return paginationInfo, nil
}

func (r *bankOfResumeRepository) GetByID(id uint) (*models.BankOfResume, error) {
	r.dbManager.DB = r.qm.DB.Preload("Subtopics").Preload("User")

	var bankOfResume models.BankOfResume
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&bankOfResume, conditions); err != nil {
		return nil, err
	}
	return &bankOfResume, nil
}

func (r *bankOfResumeRepository) Update(bankOfResume *models.BankOfResume, updates map[string]interface{}) error {
	return r.dbManager.Update(bankOfResume, updates)
}

func (r *bankOfResumeRepository) Delete(id uint) error {
	bankOfResume := models.BankOfResume{ID: id}
	return r.dbManager.Delete(&bankOfResume)
}
