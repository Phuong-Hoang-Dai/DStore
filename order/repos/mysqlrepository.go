package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/order"
	"github.com/Phuong-Hoang-Dai/DStore/order/usecase"

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
	result := m.DB.Create(&data)

	return data.Id, result.Error
}

func (m mysqlOrderRepo) GetOrderById(id int, data *order.Order) error {
	result := m.DB.Table(order.OrderTableName).Preload("Items").First(&data, id)

	return result.Error
}

func (m mysqlOrderRepo) UpdateOrder(data *order.Order) error {
	result := m.DB.Table(order.OrderTableName).Updates(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlOrderRepo) DeleteOrder(id int, data *order.Order) error {
	result := m.DB.Table(order.OrderTableName).Where("id = ?", id).Delete(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlOrderRepo) GetOrders(pdata *[]order.Order) error {
	//p.Process()
	//result := m.DB.Table(product.ProductTableName).Limit(p.Limit).Offset(p.Offset).Find(data)

	//return result.Error
	return nil
}
