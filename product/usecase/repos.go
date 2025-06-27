package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
)

type ProductRepos interface {
	CreateProduct(data *product.Product) (int, error)
	UpdateProduct(data product.Product) error
	UpdateProducts(data []product.Product) error
	GetProductById(id int) (product.Product, error)
	GetProducts(p product.Paging) ([]product.Product, error)
	DeleteProduct(id int) error
}
