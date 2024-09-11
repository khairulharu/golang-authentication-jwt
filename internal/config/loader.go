package config

import (
	"os"
)

func New() *Config {
	return &Config{
		Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
		Server{
			Host: os.Getenv("SRV_HOST"),
			Port: os.Getenv("SRV_PORT"),
		},
	}
}
