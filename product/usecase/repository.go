package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
)

type ProductRepository interface {
	CreateProduct(data *product.Product) (int, error)
	UpdateProduct(data *product.Product) error
	GetProductByid(id int, data *product.Product) error
	GetProducts(p *product.Paging, data *[]product.Product) error
	DeleteProduct(id int, data *product.Product) error
}
