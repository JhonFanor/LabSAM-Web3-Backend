package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type UniversityUserRepository interface {
	Create(universityUser *models.UniversityUser) (*models.UniversityUser, error)
	Update(universityUser *models.UniversityUser) error
	Delete(userID uint) error
	GetByUserID(userID uint) (*models.UniversityUser, error)
	GetAll() ([]models.UniversityUser, error)
}

type universityUserRepository struct {
	db *gorm.DB
}

func NewUniversityUserRepository(db *gorm.DB) UniversityUserRepository {
	return &universityUserRepository{db: db}
}

func (r *universityUserRepository) Create(universityUser *models.UniversityUser) (*models.UniversityUser, error) {
	if err := r.db.Create(universityUser).Error; err != nil {
		return nil, err
	}
	return universityUser, nil
}

func (r *universityUserRepository) Update(universityUser *models.UniversityUser) error {
	return r.db.Save(universityUser).Error
}

func (r *universityUserRepository) Delete(userID uint) error {
	return r.db.Delete(&models.UniversityUser{}, userID).Error
}

func (r *universityUserRepository) GetByUserID(userID uint) (*models.UniversityUser, error) {
	var universityUser models.UniversityUser
	if err := r.db.First(&universityUser, userID).Error; err != nil {
		return nil, err
	}
	return &universityUser, nil
}

func (r *universityUserRepository) GetAll() ([]models.UniversityUser, error) {
	var universityUsers []models.UniversityUser
	if err := r.db.Find(&universityUsers).Error; err != nil {
		return nil, err
	}
	return universityUsers, nil
}
