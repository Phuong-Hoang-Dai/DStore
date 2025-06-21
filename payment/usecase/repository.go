package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/order"
)

type OrderRepository interface {
	Payment(data *order.Order) (int, error)
	GetPaymenById(id int, data *order.Order) error
}
