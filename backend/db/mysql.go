package db

import (
	"github.com/Phuong-Hoang-Dai/DStore/configs"
	"github.com/Phuong-Hoang-Dai/DStore/internal/order"
	"github.com/Phuong-Hoang-Dai/DStore/internal/product"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() (db *gorm.DB, err error) {
	dsn := configs.Cfg.ConnectStr
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&product.Product{}, &order.Order{}, &order.OrderItem{}, &user.User{})
	if err != nil {
		return nil, err
	}
	err = SeedData(db)

	if err != nil {
		return db, err
	}

	return db, nil
}
