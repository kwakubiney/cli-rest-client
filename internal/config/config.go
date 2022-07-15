package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadNormalConfig() error {
	pathToEnvFile := "./.env"
	log.Println(pathToEnvFile)
	if _, err := os.Stat(pathToEnvFile); err == nil {
		return godotenv.Load(pathToEnvFile)
	}

	return nil
}
