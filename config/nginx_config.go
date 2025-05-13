package config

import "os"

type NginxConfig struct {
	NGINX_URL string
}

func NewNginxConfig() *NginxConfig {
	print("aqui aqui", os.Getenv("NGINX_URL"))
	return &NginxConfig{
		NGINX_URL: os.Getenv("NGINX_URL"),
	}
}
