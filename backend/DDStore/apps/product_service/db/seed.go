package db

import (
	"errors"
	"log"

	"github.com/Phuong-Hoang-Dai/DDStore/app/product_service/internal/model"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) error {

	p := model.Category{}
	if !errors.Is(db.First(&p).Error, gorm.ErrRecordNotFound) {
		log.Println("Data is seeded")
		return nil
	}

	cates := []model.Category{
		{Id: 1, Name: "Other"},
		{Id: 2, Name: "Drink"},
	}

	if err := db.Create(&cates).Error; err != nil {
		return err
	}

	products := []model.Product{
		{Name: "Cà phê sữa", Desc: "Cà phê pha với sữa đặc", Price: 25000, Quantity: 100, CategoryID: 1},
		{Name: "Trà đào", Desc: "Trà đào mát lạnh", Price: 30000, Quantity: 80, CategoryID: 1},
		{Name: "Bạc xỉu", Desc: "Sữa nhiều cà phê ít", Price: 28000, Quantity: 60, CategoryID: 2},
		{Name: "Cà phê đen", Desc: "Cà phê nguyên chất", Price: 22000, Quantity: 90, CategoryID: 1},
		{Name: "Trà sữa trân châu", Desc: "Trà sữa thơm béo", Price: 35000, Quantity: 120, CategoryID: 2},
	}

	if err := db.Create(&products).Error; err != nil {
		return err
	}

	return nil
}
