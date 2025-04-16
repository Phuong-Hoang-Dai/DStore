package product

import "time"

type Product struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Desc      string    `json:"description" gorm:"column:description"`
	Price     float32   `json:"price" gorm:"column:price"`
	Create_at time.Time `json:"create_at" gorm:"column:create_at"`
	Update_at time.Time `json:"update_at" gorm:"column:update_at"`
}

func (Product) GetTableName() string { return "Product" }
