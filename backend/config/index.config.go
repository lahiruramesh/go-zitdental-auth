package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env string
	ZitadelDomain string
	ZitadelClientSecret string
	ZitadelCLientID string
	AuthCallbackURL string
	WebCallbackURL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config {
		Port: os.Getenv("PORT"),
		Env: os.Getenv("ENV"),
		ZitadelDomain: os.Getenv("ZITADEL_DOMAIN"),
		ZitadelClientSecret: os.Getenv("ZITADETAL_CLIENT_SECRET"),
		ZitadelCLientID: os.Getenv("ZITADEL_CLIENT_ID"),
		AuthCallbackURL: os.Getenv("AUTH_CALLBACK_URL"),
		WebCallbackURL: os.Getenv("WEB_CALLBACK_URL"),
	}
}

func GetEnvConfig(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return "NO_ENV"
}