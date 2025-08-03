package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type BusinessUserRepository interface {
	Create(businessUser *models.BusinessUser) (*models.BusinessUser, error)
	Update(businessUser *models.BusinessUser, updates map[string]interface{}) error
	Delete(userID uint) error
	GetByUserID(userID uint) (*models.BusinessUser, error)
	GetAll() ([]models.BusinessUser, error)
}

type businessUserRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewBusinessUserRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) BusinessUserRepository {
	return &businessUserRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *businessUserRepository) Create(businessUser *models.BusinessUser) (*models.BusinessUser, error) {
	if err := r.db.Create(businessUser).Error; err != nil {
		return nil, err
	}
	return businessUser, nil
}

func (r *businessUserRepository) Update(businessUser *models.BusinessUser, updates map[string]interface{}) error {
	return r.dbManager.Update(businessUser, updates, r.db)
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
