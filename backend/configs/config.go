package configs

import (
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
	godotenv.Load("../configs/.env")
	Cfg = Config{
		AppName:     os.Getenv("APP_NAME"),
		ConnectStr:  os.Getenv("DB_CONN_STR"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpireIn: os.Getenv("JWT_EXPIRES_IN"),
		SysAccount:  os.Getenv("SYSTEM_ACCOUNT"),
		SysPassword: os.Getenv("SYSTEM_PW"),
	}
}
