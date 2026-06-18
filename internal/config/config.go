package config

import (
	"log"

	"github.com/lubosgarancovsky/go-kit"
)

var GlobalConfig *Config

type Config struct {
	ServiceName        string
	Port               int    `field:"PORT" default:"9092"`
	DBUrl              string `field:"DB_URL"`
	UploadsFolder      string `field:"UPLOADS_FOLDER"`
	TemplatesFolder    string `field:"TEMPLATES_FOLDER"`
	SMTPHost           string `field:"SMTP_HOST"`
	SMTPPort           int    `field:"SMTP_PORT"`
	SMTPFrom           string `field:"SMTP_FROM"`
	InvitationTokenExp int    `field:"INVITATION_TOKEN_EXP" default:"3600"`
	InvitationUrl      string `field:"INVITATION_URL"`
	OllamaVisionModel  string `field:"OLLAMA_VISION_MODEL"`
	OllamaUrl          string `field:"OLLAMA_URL"`
}

func Init(serviceName string) {
	var appConfig Config
	if err := go_kit.LoadEnv(&appConfig, ".env", ".env.local"); err != nil {
		log.Fatal("Failed to load config from .env file", err)
	}

	appConfig.ServiceName = serviceName
	GlobalConfig = &appConfig
}
