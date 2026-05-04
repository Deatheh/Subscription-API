package config

import (
	"fmt"
	"os"
	"strings"
)

type Application struct {
	Port    string
	LogPath string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Application Application
	Db          Db
}

func NewEnvConfig() *Config {
	return &Config{
		Application: Application{
			Port:    os.Getenv("APP_PORT"),
			LogPath: os.Getenv("APP_LOG_PATH"),
		},
		Db: Db{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		},
	}
}

func PrintConfigWithHiddenSecrets(config *Config) {
	mask := func(s string) string {
		if s == "" {
			return ""
		}
		return strings.Repeat("*", len(s))
	}

	fmt.Println("=== App Config ===")
	fmt.Printf("App port: %s\n", config.Application.Port)
	fmt.Printf("App log path: %s\n", config.Application.LogPath)

	fmt.Println("=== DB Config ===")
	fmt.Printf("Host: %s\n", config.Db.Host)
	fmt.Printf("Port: %s\n", config.Db.Port)
	fmt.Printf("User: %s\n", config.Db.User)
	fmt.Printf("Password: %s\n", mask(config.Db.Password))
	fmt.Printf("Database: %s\n", config.Db.Database)
}
