package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type BankOfResumeSubtopicRepository interface {
	Create(bankOfResumeSubtopic *models.BankOfResumeSubtopic) (*models.BankOfResumeSubtopic, error)
	Delete(bankOfResumeID, subtopicID uint) error
}

type bankOfResumeSubtopicRepository struct {
	db *gorm.DB
}

func NewBankOfResumeSubtopicRepository(db *gorm.DB) BankOfResumeSubtopicRepository {
	return &bankOfResumeSubtopicRepository{
		db: db,
	}
}

func (r *bankOfResumeSubtopicRepository) Create(bankOfResumeSubtopic *models.BankOfResumeSubtopic) (*models.BankOfResumeSubtopic, error) {
	if err := r.db.Create(bankOfResumeSubtopic).Error; err != nil {
		return nil, err
	}
	return bankOfResumeSubtopic, nil
}

func (r *bankOfResumeSubtopicRepository) Delete(bankOfResumeID, subtopicID uint) error {
	return r.db.Where("bank_of_resume_id = ? AND subtopic_id = ?", bankOfResumeID, subtopicID).
		Delete(&models.BankOfResumeSubtopic{}).Error
}
