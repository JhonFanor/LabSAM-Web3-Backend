package config

import "os"

type ServerConfig struct {
	SERVER   string
	FRONTEND string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		SERVER:   os.Getenv("SERVER"),
		FRONTEND: os.Getenv("FRONTEND"),
	}
}
