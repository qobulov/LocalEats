package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE  string
	USER_ROUTER   string
	ORDER_SERVICE string
	DB_HOST       string
	DB_PORT       string
	DB_USER       string
	DB_PASSWORD   string
	DB_NAME       string
	SIGNING_KEY   string
	GMAIL_CODE    string
	EXTRA_SERVICE string
}

func Load() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current path: ", path)


	if err := godotenv.Load(path+"/.env"); err != nil {
		log.Print("No .env file found")
	}


	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "hamidjon4424"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "restaurants"))
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", "50051"))
	config.USER_ROUTER = cast.ToString(Coalesce("USER_ROUTER", "50051"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))
	config.GMAIL_CODE = cast.ToString(Coalesce("GMAIL_CODE", "secret"))
	config.ORDER_SERVICE = cast.ToString(Coalesce("ORDER_SERVICE", "6666"))
	config.EXTRA_SERVICE = cast.ToString(Coalesce("EXTRA_SERVICE", "5555"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
