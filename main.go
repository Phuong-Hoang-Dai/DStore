package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Phuong-Hoang-Dai/DStore/handler"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}
	fmt.Println(db)

	handler.SetupHttp()
}
