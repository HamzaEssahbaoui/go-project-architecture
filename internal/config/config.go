package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

func Load(filepath string) (*Config, error) {
	err := godotenv.Load(filepath)
	if err != nil {
		apiKey := os.Getenv("TMDB_API_KEY")
		if apiKey == "" {
			return nil, errors.New("TMDB_API_KEY is not set")
		}
		return nil, err
	}

	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, err
	}

	return &Config{APIKey: apiKey}, nil
}
