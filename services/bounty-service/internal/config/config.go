package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig(filenames ...string) (*Config, error) {
	_ = godotenv.Load(filenames...)

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}

	if cfg.DatabaseURL == "" {
		return nil, errors.New("DATABASE_URL is not set in environment variables")
	}
	if cfg.ServerPort == "" {
		return nil, errors.New("SERVER_PORT is not set in environment variables")
	}

	return cfg, nil
}
