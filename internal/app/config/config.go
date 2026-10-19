package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/angelusora/catalog-service/internal/app/config/section"
)

type Config struct {
	Repository section.Repository
	Processor  section.Processor
	Monitor    section.Monitor
}

var Root Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Не удалось загрузить .en: %v", err)
	}

	if err := envconfig.Process("APP", &Root); err != nil {
		log.Fatal(err)
	}
}
