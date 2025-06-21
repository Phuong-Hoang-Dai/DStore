package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id          int            `json:"orderId" gorm:"column:orderId"`
	State       int            `json:"state" gorm:"column:state"`
	Items       []OrderItem    `json:"items" gorm:"foreignKey:OrderId;references:Id"`
	Total       int            `json:"total" gorm:"column:total"`
	CreateAt    time.Time      `json:"create_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at" gorm:"index"`
	IsValidated bool
}

type OrderItem struct {
	ProductId      int  `json:"productId" gorm:"column:productId"`
	Quantity       int  `json:"quantity" gorm:"column:quantity"`
	OrderId        int  `json:"orderId" gorm:"column:orderId"`
	IsConfirmed    bool `json:"is_conformed"`
	IsUpdatedStock bool `json:"is_updated_stock"`
}

func (order Order) Validate() bool {
	if order.Items == nil {
		return false
	}

	for _, v := range order.Items {
		if !v.IsConfirmed || !v.IsUpdatedStock {
			return false
		}
	}

	return order.Total != 0
}

const (
	Pending = iota
	IsPaid
	Completed
	Cancelled
)
