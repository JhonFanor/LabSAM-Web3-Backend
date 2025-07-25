package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) (*models.Contact, error)
	Update(contact *models.Contact, updates map[string]interface{}) error
}

type contactRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
}

func NewContactRepository(dbManager *gormmanagers.DBManager, db *gorm.DB) ContactRepository {
	return &contactRepository{
		dbManager: dbManager,
		db:        db,
	}
}

func (r *contactRepository) Create(contact *models.Contact) (*models.Contact, error) {
	if err := r.dbManager.Create(contact, r.db); err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *contactRepository) Update(contact *models.Contact, updates map[string]interface{}) error {
	return r.dbManager.Update(contact, updates, r.db)
}
