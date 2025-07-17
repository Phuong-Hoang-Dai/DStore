package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/internal/product"
)

type MockRepos struct{}

var productList []product.Product

func (MockRepos) Init() {
	productList = []product.Product{}
}

func (MockRepos) CreateProduct(data *product.Product) (int, error) {
	productList = append(productList, *data)
	productList[len(productList)-1].Id = len(productList)

	return productList[len(productList)-1].Id, nil
}

func (MockRepos) UpdateProduct(data product.Product) error {
	productList[data.Id-1] = data

	return nil
}

func (MockRepos) UpdateProducts(data []product.Product) error {
	for i := range data {
		productList[data[i].Id-1] = data[i]
	}

	return nil
}

func (MockRepos) GetProductById(id int) (product.Product, error) {
	return productList[id-1], nil
}

func (MockRepos) GetProducts(p product.Paging) (data []product.Product, err error) {
	for i := p.Offset; i < len(productList); i++ {
		if i-p.Offset < p.Limit {
			data = append(data, productList[i])
		}
	}

	return data, nil
}

func (MockRepos) DeleteProduct(id int) error {
	return nil
}
