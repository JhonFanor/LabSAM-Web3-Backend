package config

import "os"

type FileServerConfig struct {
	FILE_SERVER_URL string
}

func NewFileServerConfig() *FileServerConfig {
	return &FileServerConfig{
		FILE_SERVER_URL: os.Getenv("FILE_SERVER_URL"),
	}
}
