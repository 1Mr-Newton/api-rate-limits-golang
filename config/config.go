package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Port                int
	GlobalRateLimit     int
	UserRateLimit       int
	AdminRateLimit      int
	RateLimitExpiration int
}

// GetAddress returns the formatted address string for the server
func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	return &Config{
		Port:                getEnvAsInt("PORT", 3000),
		GlobalRateLimit:     getEnvAsInt("GLOBAL_RATE_LIMIT", 100),    // 100 requests per minute by default
		UserRateLimit:       getEnvAsInt("USER_RATE_LIMIT", 50),       // 50 requests per minute by default
		AdminRateLimit:      getEnvAsInt("ADMIN_RATE_LIMIT", 200),     // 200 requests per minute by default
		RateLimitExpiration: getEnvAsInt("RATE_LIMIT_EXPIRATION", 60), // 60 seconds (1 minute) by default
	}
}

// getEnvAsInt gets an environment variable as an integer with a fallback value
func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
