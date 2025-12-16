package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration settings.
type Config struct {
	AppEnv         string        `mapstructure:"APP_ENV"`
	AppPort        int           `mapstructure:"APP_PORT"`
	DatabaseURL    string        `mapstructure:"DATABASE_URL"`
	CacheRedisAddr string        `mapstructure:"CACHE_REDIS_ADDR"`
	AuthSecret     string        `mapstructure:"AUTH_SECRET"`
	LogLevel       string        `mapstructure:"LOG_LEVEL"`
	RateLimit      int           `mapstructure:"RATE_LIMIT_PER_MINUTE"`
	Timeout        time.Duration `mapstructure:"TIMEOUT_SECONDS"`
}

// LoadConfig loads configuration from environment variables or a config file.
func LoadConfig() (config Config, err error) {
	// Set sane defaults
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("RATE_LIMIT_PER_MINUTE", 60)
	viper.SetDefault("TIMEOUT_SECONDS", 10)

	// Environment variables
	viper.AutomaticEnv()

	// If a config file is present (e.g., .env)
	viper.AddConfigPath(".") // Look for config file in current directory
	viper.SetConfigName("app") // Name of config file (without extension)
	viper.SetConfigType("env") // Type of config file (e.g., "json", "yaml", "env")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using environment variables and defaults.")
		} else {
			return // Some other error occurred
		}
	}

	// Unmarshal the config into the struct
	err = viper.Unmarshal(&config)

	// Special handling for time.Duration from env var
	if timeoutStr := os.Getenv("TIMEOUT_SECONDS"); timeoutStr != "" {
		if timeoutInt, convErr := strconv.Atoi(timeoutStr); convErr == nil {
			config.Timeout = time.Duration(timeoutInt) * time.Second
		} else {
			log.Printf("Warning: Could not parse TIMEOUT_SECONDS '%s', using default. Error: %v", timeoutStr, convErr)
		}
	} else if config.Timeout == 0 { // if not set by file or env, use default (viper.Unmarshal might not respect time.Duration directly)
		config.Timeout = viper.GetDuration("TIMEOUT_SECONDS") * time.Second
	}

	return
}
