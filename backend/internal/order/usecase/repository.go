package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/internal/order"
)

type OrderRepository interface {
	CreateOrder(data *order.Order) (int, error)
	UpdateOrder(data order.Order) error
	GetOrderById(id int) (order.Order, error)
	GetOrders(p order.Paging) ([]order.Order, error)
	GetHistoryOrders(id int, p order.Paging) ([]order.Order, error)
	DeleteOrder(id int) error
}
