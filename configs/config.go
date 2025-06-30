package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	ConnectStr  string
	JWTSecret   string
	JWTExpireIn string
	SysAccount  string
	SysPassword string
}

var Cfg Config

func LoadConfig() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Cfg = Config{
		AppName:     os.Getenv("APP_NAME"),
		ConnectStr:  os.Getenv("DB_CONN_STR"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpireIn: os.Getenv("JWT_EXPIRES_IN"),
		SysAccount:  os.Getenv("SYSTEM_ACCOUNT"),
		SysPassword: os.Getenv("SYSTEM_PW"),
	}
}
