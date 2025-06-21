package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/order"
)

type OrderRepository interface {
	CreateOrder(data *order.Order) (int, error)
	UpdateOrder(data *order.Order) error
	GetOrderById(id int, data *order.Order) error
	GetOrders(data *[]order.Order) error
	DeleteOrder(id int, data *order.Order) error
}
