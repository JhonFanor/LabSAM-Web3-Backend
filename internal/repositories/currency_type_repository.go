package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type CurrencyTypeRepository interface {
	Create(currencyType *models.CurrencyType) (*models.CurrencyType, error)
	GetByCode(code string) (*models.CurrencyType, error)
}

type currencyTypeRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewCurrencyTypeRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) *currencyTypeRepository {
	return &currencyTypeRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *currencyTypeRepository) Create(currencyType *models.CurrencyType) (*models.CurrencyType, error) {
	if err := r.dbManager.Create(currencyType, r.db); err != nil {
		return nil, err
	}
	return currencyType, nil
}

func (r *currencyTypeRepository) GetByCode(code string) (*models.CurrencyType, error) {
	var currencyType models.CurrencyType
	if err := r.db.
		Where("code = ?", code).
		First(&currencyType).Error; err != nil {
		return nil, err
	}
	return &currencyType, nil
}
