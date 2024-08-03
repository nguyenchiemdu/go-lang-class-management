package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port       string `json:"port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbPort     string `json:"db_port"`
	DbName     string `json:"db_name"`
	JWTSecret  string `json:"jwt_secret"`
	JWTExpire  int    `json:"jwt_expire"`
}

func LoadAppConfig() AppConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return AppConfig{}

	}

	JWTExpire, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))

	if err != nil {
		log.Fatal("Error parsing JWT_EXPIRE")
	}

	return AppConfig{
		Port:       os.Getenv("PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTExpire:  JWTExpire,
	}
}
