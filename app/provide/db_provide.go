package provide

import (
	"fmt"
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseProvide(logger logging.Logger) (*gorm.DB, error) {
	config := config.NewDatabaseConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogError("Failed to connect to the database", err)
		return nil, err
	}

	logger.LogInfo("Successfully connected to the database")

	return db, nil
}
