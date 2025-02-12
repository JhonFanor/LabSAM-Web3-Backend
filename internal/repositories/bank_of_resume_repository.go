package repositories

import (
	"context"
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type BankOfResumeRepository interface {
	Create(bankOfResume *models.BankOfResume) (*models.BankOfResume, error)
	Update(bankOfResume *models.BankOfResume) error
	Delete(id uint) error
	GetByID(id uint) (*models.BankOfResume, error)
	GetAll(ctx context.Context) ([]models.BankOfResume, *dto.PaginationDTO, error)
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

func (r *bankOfResumeRepository) Update(bankOfResume *models.BankOfResume) error {
	return r.db.Save(bankOfResume).Error
}

func (r *bankOfResumeRepository) Delete(id uint) error {
	return r.db.Delete(&models.BankOfResume{}, id).Error
}

func (r *bankOfResumeRepository) GetByID(id uint) (*models.BankOfResume, error) {
	var bankOfResume models.BankOfResume
	if err := r.db.First(&bankOfResume, id).Error; err != nil {
		return nil, err
	}
	return &bankOfResume, nil
}

func (r *bankOfResumeRepository) GetAll(ctx context.Context) ([]models.BankOfResume, *dto.PaginationDTO, error) {
	var bankOfResumes []models.BankOfResume

	tx := r.db.Model(&models.BankOfResume{})

	tx, paginationInfo, err := r.qm.ApplyPaginationAndFilters(ctx, tx, &models.BankOfResume{})
	if err != nil {
		return nil, nil, err
	}

	if err := tx.Find(&bankOfResumes).Error; err != nil {
		return nil, nil, err
	}

	return bankOfResumes, paginationInfo, nil
}
