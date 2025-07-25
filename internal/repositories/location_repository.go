package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(location *models.Location) (*models.Location, error)
	GetByCountryAndCity(country string, city string) (*models.Location, error)
}

type locationRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewLocationRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) LocationRepository {
	return &locationRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *locationRepository) Create(location *models.Location) (*models.Location, error) {
	if err := r.dbManager.Create(location, r.db); err != nil {
		return nil, err
	}
	return location, nil
}

func (r *locationRepository) GetByCountryAndCity(country string, city string) (*models.Location, error) {
	var location models.Location
	if err := r.db.
		Where("country = ? AND city = ?", country, city).
		First(&location).Error; err != nil {
		return nil, err
	}
	return &location, nil
}
