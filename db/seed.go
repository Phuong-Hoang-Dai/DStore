package db

import (
	"errors"
	"log"

	"github.com/Phuong-Hoang-Dai/DStore/order"
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) error {
	p := product.Product{}
	if !errors.Is(db.First(&p).Error, gorm.ErrRecordNotFound) {
		log.Println("Data is seeded")
		return nil
	}

	products := []product.Product{
		{Name: "Cà phê sữa", Desc: "Cà phê pha với sữa đặc", Price: 25000, Quantity: 100},
		{Name: "Trà đào", Desc: "Trà đào mát lạnh", Price: 30000, Quantity: 80},
		{Name: "Bạc xỉu", Desc: "Sữa nhiều cà phê ít", Price: 28000, Quantity: 60},
		{Name: "Cà phê đen", Desc: "Cà phê nguyên chất", Price: 22000, Quantity: 90},
		{Name: "Trà sữa trân châu", Desc: "Trà sữa thơm béo", Price: 35000, Quantity: 120},
	}

	if err := db.Create(&products).Error; err != nil {
		return err
	}

	orders := []order.Order{
		{State: 0, Total: 0},
		{State: 0, Total: 0},
	}

	if err := db.Create(&orders).Error; err != nil {
		return err
	}

	orderItems := []order.OrderItem{
		{ProductId: products[0].Id, Quantity: 2, OrderId: orders[0].Id},
		{ProductId: products[1].Id, Quantity: 1, OrderId: orders[0].Id},

		{ProductId: products[2].Id, Quantity: 1, OrderId: orders[1].Id},
		{ProductId: products[4].Id, Quantity: 2, OrderId: orders[1].Id},
	}

	if err := db.Create(&orderItems).Error; err != nil {
		return err
	}

	return nil
}
