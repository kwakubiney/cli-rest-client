package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadTestConfig() error {
	pathToTestEnvFile := "../../.env.test"
	if _, err := os.Stat(pathToTestEnvFile); err == nil {
		return godotenv.Load(pathToTestEnvFile)
	}
	return nil
}

func LoadNormalConfig() error {
	pathToEnvFile := "./.env"
	if _, err := os.Stat(pathToEnvFile); err == nil {
		return godotenv.Load(pathToEnvFile)
	}

	return nil
}
