package gormmanagers

import (
	"errors"
	"lamsam-web3-backend/internal/logging"

	"gorm.io/gorm"
)

type DBManager struct {
	DB           *gorm.DB
	QueryManager *GormQueryManager
	Logger       logging.Logger
}

func NewDBManager(db *gorm.DB, queryManager *GormQueryManager, logger logging.Logger) *DBManager {
	return &DBManager{
		DB:           db,
		QueryManager: queryManager,
		Logger:       logger,
	}
}

func (m *DBManager) Create(model interface{}) error {
	if err := m.DB.Create(model).Error; err != nil {
		m.Logger.LogError("Error creating record", err)
		return err
	}
	m.Logger.LogInfo("Record created successfully")
	return nil
}

func (m *DBManager) Find(model interface{}, conditions map[string]interface{}) error {
	if err := m.DB.Where(conditions).First(model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			m.Logger.LogWarn("Record not found", conditions)
		}
		m.Logger.LogError("Error finding record", err)
		return err
	}
	m.Logger.LogInfo("Record found successfully", conditions)
	return nil
}

func (m *DBManager) Update(model interface{}, updates map[string]interface{}) error {
	if err := m.DB.Model(model).Updates(updates).Error; err != nil {
		m.Logger.LogError("Error updating record", err)
		return err
	}
	m.Logger.LogInfo("Record updated successfully", updates)
	return nil
}

func (m *DBManager) Delete(model interface{}) error {
	if err := m.DB.Delete(model).Error; err != nil {
		m.Logger.LogError("Error deleting record", err)
		return err
	}
	m.Logger.LogInfo("Record deleted successfully")
	return nil
}
