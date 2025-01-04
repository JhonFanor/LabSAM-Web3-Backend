package gormadapter

import (
	"LABSAM-WEB3-BACKEND/config"
	"LABSAM-WEB3-BACKEND/internal/database"
	"LABSAM-WEB3-BACKEND/internal/logging"

	"gorm.io/gorm"
)

type GormDBManager struct {
	db           *gorm.DB
	queryManager *GormQueryManager
	l            logging.Logger
	appConfig    *config.AppConfig
}

func NewGormDBManager(db *gorm.DB, queryManager *GormQueryManager, logger logging.Logger, appConfig *config.AppConfig) database.DBManager {
	return &GormDBManager{
		db:           db,
		queryManager: queryManager,
		l:            logger,
		appConfig:    appConfig,
	}
}

func (g *GormDBManager) Migrate() error {
	g.l.LogInfo("Running migrations...")
	// Implementación de migraciones
	return nil
}

func (g *GormDBManager) PreloadMasterData() error {
	g.l.LogInfo("Preloading master data...")
	// Implementación para precargar datos maestros
	return nil
}
