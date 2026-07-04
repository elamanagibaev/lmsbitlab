package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

func Load() *Config {
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("POSTGRES_HOST", "localhost"),
		DBPort:     getEnv("POSTGRES_PORT", "5434"),
		DBUser:     getEnv("POSTGRES_USER", "lmsbitlab"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:     getEnv("POSTGRES_DB", "lmsbitlab"),
		DBSSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
	}
}
