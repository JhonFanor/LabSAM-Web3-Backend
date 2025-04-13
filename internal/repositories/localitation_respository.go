package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"
)

type LocalitationRepository interface {
	Create(localitation *models.Localitation) (*models.Localitation, error)
	GetByFields(localitation *models.Localitation) (*models.Localitation, error)
	IsInUse(id uint) (bool, error)
	Delete(id uint) error
}

type localitationRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
}

func NewLocalitationRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) LocalitationRepository {
	return &localitationRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *localitationRepository) Create(localitation *models.Localitation) (*models.Localitation, error) {
	if err := r.dbManager.Create(localitation); err != nil {
		return nil, err
	}
	return localitation, nil
}

func (r *localitationRepository) GetByFields(localitation *models.Localitation) (*models.Localitation, error) {

	conditions := map[string]interface{}{
		"address":   localitation.Address,
		"latitude":  localitation.Latitude,
		"longitude": localitation.Longitude,
	}

	if err := r.dbManager.Find(&localitation, conditions); err != nil {
		return nil, err
	}
	return localitation, nil
}

func (r *localitationRepository) IsInUse(id uint) (bool, error) {

	var count int64

	tables := []string{"events", "companies"}
	for _, table := range tables {
		err := r.dbManager.DB.Table(table).Where("localitation_id = ?", id).Count(&count).Error
		if err != nil {
			return false, err
		}
		if count > 0 {
			return true, nil
		}
	}

	return false, nil
}

func (r *localitationRepository) Delete(id uint) error {
	inUse, err := r.IsInUse(id)
	if err != nil {
		return err
	}
	if inUse {
		return nil
	}
	localitation := models.Localitation{ID: id}
	if err := r.dbManager.Delete(&localitation); err != nil {
		return err
	}
	return nil
}
