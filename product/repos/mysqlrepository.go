package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/Phuong-Hoang-Dai/DStore/product/usecase"

	"gorm.io/gorm"
)

type mysqlProductRepo struct {
	DB *gorm.DB
}

func NewMysqlProductRepo(db *gorm.DB) usecase.ProductRepos {
	return mysqlProductRepo{
		DB: db,
	}
}

func (m mysqlProductRepo) CreateProduct(data *product.Product) (int, error) {
	result := m.DB.Create(&data)

	return data.Id, result.Error
}

func (m mysqlProductRepo) GetProductById(id int) (product.Product, error) {
	data := product.Product{Id: id}
	result := m.DB.First(&data)

	return data, result.Error
}

func (m mysqlProductRepo) UpdateProduct(data product.Product) error {
	result := m.DB.Updates(data)

	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (m mysqlProductRepo) UpdateProducts(data []product.Product) error {
	return m.DB.Transaction(func(tx *gorm.DB) error {
		for i := range data {
			result := tx.Updates(data[i])
			if result.RowsAffected == 0 && result.Error == nil {
				return gorm.ErrRecordNotFound
			}
		}
		return nil
	})
}

func (m mysqlProductRepo) DeleteProduct(id int) error {
	result := m.DB.Delete(&product.Product{Id: id})
	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (m mysqlProductRepo) GetProducts(p product.Paging) ([]product.Product, error) {
	data := []product.Product{}
	result := m.DB.Limit(p.Limit).Offset(p.Offset).Find(&data)

	return data, result.Error
}
