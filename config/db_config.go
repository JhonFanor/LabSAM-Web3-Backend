package config

import (
	"lamsam-web3-backend/internal/utils"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

func NewDatabaseConfig() *DatabaseConfig {
	log.Print("port de la base de datos", os.Getenv("DB_PORT"))
	log.Print("port de la base de datos", os.Getenv("JWT_SECRETE_KEY"))
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     utils.MustAtoi(os.Getenv("DB_PORT")),
	}
}
