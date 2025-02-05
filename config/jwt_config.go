package config

import "os"

type JwtConfig struct {
	SECRET_KEY     []byte
	REFRESH_SECRET []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		SECRET_KEY:     []byte(os.Getenv("JWT_SECRETE_KEY")),
		REFRESH_SECRET: []byte(os.Getenv("JWT_REFRESH_SECRET")),
	}
}
