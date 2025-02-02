package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type BusinessUserRepository interface {
	Create(businessUser *models.BusinessUser) (*models.BusinessUser, error)
	Update(businessUser *models.BusinessUser) error
	Delete(userID uint) error
	GetByUserID(userID uint) (*models.BusinessUser, error)
	GetAll() ([]models.BusinessUser, error)
}

type businessUserRepository struct {
	db *gorm.DB
}

func NewBusinessUserRepository(db *gorm.DB) BusinessUserRepository {
	return &businessUserRepository{db: db}
}

func (r *businessUserRepository) Create(businessUser *models.BusinessUser) (*models.BusinessUser, error) {
	if err := r.db.Create(businessUser).Error; err != nil {
		return nil, err
	}
	return businessUser, nil
}

func (r *businessUserRepository) Update(businessUser *models.BusinessUser) error {
	return r.db.Save(businessUser).Error
}

func (r *businessUserRepository) Delete(userID uint) error {
	return r.db.Delete(&models.BusinessUser{}, userID).Error
}

func (r *businessUserRepository) GetByUserID(userID uint) (*models.BusinessUser, error) {
	var businessUser models.BusinessUser
	if err := r.db.First(&businessUser, userID).Error; err != nil {
		return nil, err
	}
	return &businessUser, nil
}

func (r *businessUserRepository) GetAll() ([]models.BusinessUser, error) {
	var businessUsers []models.BusinessUser
	if err := r.db.Find(&businessUsers).Error; err != nil {
		return nil, err
	}
	return businessUsers, nil
}
