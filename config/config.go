package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database ConfigDatabase
	Server   ConfigServer
}

type ConfigServer struct {
	Port string
	Mode string
}

type ConfigDatabase struct {
	DSN string
}

func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
	return &Config{
		Database: ConfigDatabase{
			DSN: os.Getenv("DB_DSN"),
		},
		Server: ConfigServer{
			Port: os.Getenv("SERVER_PORT"),
			Mode: os.Getenv("SERVER_MODE"),
		},
	}
}
