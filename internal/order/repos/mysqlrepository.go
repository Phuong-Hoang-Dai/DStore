package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/internal/order"
	"github.com/Phuong-Hoang-Dai/DStore/internal/order/usecase"

	"gorm.io/gorm"
)

type mysqlOrderRepo struct {
	DB *gorm.DB
}

func NewMysqlOrderRepo(db *gorm.DB) usecase.OrderRepository {
	return mysqlOrderRepo{
		DB: db,
	}
}

func (m mysqlOrderRepo) CreateOrder(data *order.Order) (int, error) {
	result := m.DB.Create(data)

	return data.Id, result.Error
}

func (m mysqlOrderRepo) GetOrderById(id int) (data order.Order, err error) {
	result := m.DB.Preload("Items").First(&data, id)

	return data, result.Error
}

func (m mysqlOrderRepo) UpdateOrder(data order.Order) error {
	result := m.DB.Updates(data)
	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlOrderRepo) DeleteOrder(id int) error {
	result := m.DB.Delete(&order.Order{Id: id})
	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlOrderRepo) GetOrders(p order.Paging) (orders []order.Order, err error) {
	result := m.DB.Limit(p.Limit).Offset(p.Offset).Find(&orders)

	return orders, result.Error
}

func (m mysqlOrderRepo) GetHistoryOrders(id int, p order.Paging) (orders []order.Order, err error) {
	result := m.DB.Order("created_at DESC").Limit(p.Limit).Offset(p.Offset).Find(&orders)

	return orders, result.Error
}
