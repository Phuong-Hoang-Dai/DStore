package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"column:name"`
	Desc      string         `json:"description" gorm:"column:description"`
	Price     float32        `json:"price" gorm:"column:price"`
	Quantity  int            `json:"quantity" gorm:"column:quantity"`
	OrderItem []OrderItem    `json:"items" gorm:"foreignKey:ProductId;references:Id"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

type OrderItem struct {
	ProductId      int  `json:"productId" gorm:"column:productId"`
	Quantity       int  `json:"quantity" gorm:"column:quantity"`
	OrderId        int  `json:"orderId" gorm:"column:orderId"`
	IsConfirmed    bool `json:"is_conformed"`
	IsUpdatedStock bool `json:"is_updated_stock"`
}

type Paging struct {
	Limit  int
	Offset int
}

func (p *Paging) Process() {
	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
	if p.Limit < 0 {
		p.Limit = 0
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
}

func (Product) GetTableName() string { return "Product" }
