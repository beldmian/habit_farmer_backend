package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BindAddr string
	DbUrl string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		BindAddr: getEnv("BIND_ADDR", ":8080"),
		DbUrl: getEnv("DB_URL", ""),
	}, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}