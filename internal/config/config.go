package config

import (
	"log"

	"github.com/lubosgarancovsky/go-kit/cfg"
)

type Config struct {
	Port          int    `field:"PORT" default:"9091"`
	DBUrl         string `field:"DB_URL"`
	UploadsFolder string `field:"UPLOADS_FOLDER"`
}

func LoadConfig() *Config {
	var appConfig Config
	if err := cfg.LoadEnv(&appConfig, ".env", ".env.local"); err != nil {
		log.Fatal("Failed to load config from .env file", err)
	}

	return &appConfig
}
