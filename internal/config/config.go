package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RecaptchaSecret  string
	RecaptchaSiteKey string
	Port             string
}

func Load() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Read environment variables
	cfg := &Config{
		RecaptchaSecret:  os.Getenv("RECAPTCHA_SECRET"),
		RecaptchaSiteKey: os.Getenv("RECAPTCHA_SITE_KEY"),
		Port:             os.Getenv("PORT"),
	}

	// Validate required fields
	if cfg.RecaptchaSecret == "" {
		return nil, fmt.Errorf("RECAPTCHA_SECRET is required")
	}
	if cfg.RecaptchaSiteKey == "" {
		return nil, fmt.Errorf("RECAPTCHA_SITE_KEY is required")
	}
	if cfg.Port == "" {
		cfg.Port = "8080" // Default port
	}

	return cfg, nil
}
