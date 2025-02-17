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
	GetByID(id uint) (*models.BankOfResume, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	Update(bankOfResume *models.BankOfResume) error
	Delete(id uint) error
}

type bankOfResumeRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewBankOfResumeRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) BankOfResumeRepository {
	return &bankOfResumeRepository{
		db: db,
		qm: qm,
	}
}

func (r *bankOfResumeRepository) Create(bankOfResume *models.BankOfResume) (*models.BankOfResume, error) {
	if err := r.db.Create(bankOfResume).Error; err != nil {
		return nil, err
	}
	return bankOfResume, nil
}

func (r *bankOfResumeRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.BankOfResume{})

	return paginationInfo, nil
}

func (r *bankOfResumeRepository) GetByID(id uint) (*models.BankOfResume, error) {
	var bankOfResume models.BankOfResume
	if err := r.db.First(&bankOfResume, id).Error; err != nil {
		return nil, err
	}
	return &bankOfResume, nil
}

func (r *bankOfResumeRepository) Update(bankOfResume *models.BankOfResume) error {
	return r.db.Save(bankOfResume).Error
}

func (r *bankOfResumeRepository) Delete(id uint) error {
	return r.db.Delete(&models.BankOfResume{}, id).Error
}
