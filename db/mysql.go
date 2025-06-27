package db

import (
	"log"
	"os"

	"github.com/Phuong-Hoang-Dai/DStore/order"
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&product.Product{}, &order.Order{}, &order.OrderItem{})
	err = SeedData(db)

	if err != nil {
		return db, err
	}

	return db, nil
}
