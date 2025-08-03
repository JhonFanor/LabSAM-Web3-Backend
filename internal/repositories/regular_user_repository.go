package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type RegularUserRepository interface {
	Create(regularUser *models.RegularUser) (*models.RegularUser, error)
	Update(regularUser *models.RegularUser, updates map[string]interface{}) error
	Delete(userID uint) error
	GetByUserID(userID uint) (*models.RegularUser, error)
	GetAll() ([]models.RegularUser, error)
}

type regularUserRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewRegularUserRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) RegularUserRepository {
	return &regularUserRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *regularUserRepository) Create(regularUser *models.RegularUser) (*models.RegularUser, error) {
	if err := r.dbManager.Create(regularUser, r.db); err != nil {
		return nil, err
	}
	return regularUser, nil
}

func (r *regularUserRepository) Update(regularUser *models.RegularUser, updates map[string]interface{}) error {
	return r.dbManager.Update(regularUser, updates, r.db)
}

func (r *regularUserRepository) Delete(userID uint) error {
	return r.db.Delete(&models.RegularUser{}, userID).Error
}

func (r *regularUserRepository) GetByUserID(userID uint) (*models.RegularUser, error) {
	var regularUser models.RegularUser
	if err := r.db.First(&regularUser, userID).Error; err != nil {
		return nil, err
	}
	return &regularUser, nil
}

func (r *regularUserRepository) GetAll() ([]models.RegularUser, error) {
	var regularUsers []models.RegularUser
	if err := r.db.Find(&regularUsers).Error; err != nil {
		return nil, err
	}
	return regularUsers, nil
}
