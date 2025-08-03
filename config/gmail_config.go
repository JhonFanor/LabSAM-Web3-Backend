package config

import "os"

type GmailConfig struct {
	GMAIL_USER string
	GMAIL_PASS string
}

func NewGmailConfig() *GmailConfig {
	return &GmailConfig{
		GMAIL_USER: os.Getenv("GMAIL_USER"),
		GMAIL_PASS: os.Getenv("GMAIL_PASS"),
	}
}
