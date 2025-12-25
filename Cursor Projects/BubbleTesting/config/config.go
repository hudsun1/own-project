package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	BubbleAPIKey  string
	BubbleAppName string
	BubbleBaseURL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		BubbleAPIKey:  os.Getenv("BUBBLE_API_KEY"),
		BubbleAppName: os.Getenv("BUBBLE_APP_NAME"),
		BubbleBaseURL: os.Getenv("BUBBLE_BASE_URL"),
	}

	// Validate required configuration
	if config.BubbleAPIKey == "" {
		return nil, fmt.Errorf("BUBBLE_API_KEY is required")
	}
	if config.BubbleAppName == "" {
		return nil, fmt.Errorf("BUBBLE_APP_NAME is required")
	}
	if config.BubbleBaseURL == "" {
		config.BubbleBaseURL = "https://api.bubble.io/v1"
	}

	return config, nil
}
