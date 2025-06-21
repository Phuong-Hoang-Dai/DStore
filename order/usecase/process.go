package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/order"
)

func CreateOrder(data *order.Order, repos OrderRepository) (int, error) {
	data.State = order.Pending
	id, err := repos.CreateOrder(data)

	return id, err
}
