package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/Phuong-Hoang-Dai/DStore/product/usecase"

	"gorm.io/gorm"
)

type mysqlProductRepo struct {
	DB *gorm.DB
}

func NewMysqlProductRepo(db *gorm.DB) usecase.ProductRepository {
	return mysqlProductRepo{
		DB: db,
	}
}

func (m mysqlProductRepo) CreateProduct(data *product.Product) (int, error) {
	result := m.DB.Table(product.ProductTableName).Create(&data)

	return data.Id, result.Error
}

func (m mysqlProductRepo) GetProductByid(id int, data *product.Product) error {
	result := m.DB.Table(product.ProductTableName).First(&data, id)

	return result.Error
}

func (m mysqlProductRepo) UpdateProduct(data *product.Product) error {
	result := m.DB.Table(product.ProductTableName).Updates(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlProductRepo) DeleteProduct(id int, data *product.Product) error {
	result := m.DB.Table(product.ProductTableName).Where("id = ?", id).Delete(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlProductRepo) GetProducts(p *product.Paging, data *[]product.Product) error {
	p.Process()
	result := m.DB.Table(product.ProductTableName).Limit(p.Limit).Offset(p.Offset).Find(data)

	return result.Error
}
