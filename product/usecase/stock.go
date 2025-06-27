package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
)

func CreateProduct(data product.Product, repos ProductRepos) (int, error) {
	if id, err := repos.CreateProduct(&data); err != nil {
		return 0, err
	} else {
		return id, err
	}
}

func UpdateProduct(data product.Product, repos ProductRepos) error {
	if err := repos.UpdateProduct(data); err != nil {
		return err
	} else {
		return nil
	}
}

func GetProducts(p *product.Paging, repos ProductRepos) (data []product.Product, err error) {
	p.Process()
	if data, err = repos.GetProducts(*p); err != nil {
		return nil, err
	}
	return data, nil
}

func GetProductById(id int, repos ProductRepos) (data product.Product, err error) {
	if data, err := repos.GetProductById(id); err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func DeleteProduct(id int, repos ProductRepos) error {
	if err := repos.DeleteProduct(id); err != nil {
		return err
	} else {
		return nil
	}
}

func GetStock(data []OrderItemsDto, repos ProductRepos) error {
	stock, err := getProducts(data, repos)
	if err != nil {
		return err
	}

	for i := range data {
		if canGetStock := data[i].Quantity <= stock[i].Quantity; canGetStock {
			stock[i].Quantity -= data[i].Quantity
		} else {
			return product.ErrOutOfStock
		}
	}

	if err := repos.UpdateProducts(stock); err != nil {
		return err
	}

	return nil
}

func RestoreStock(data []OrderItemsDto, repos ProductRepos) error {
	stock, err := getProducts(data, repos)
	if err != nil {
		return err
	}

	for i := range data {
		stock[i].Quantity += (data)[i].Quantity
	}

	if err := repos.UpdateProducts(stock); err != nil {
		return err
	}

	return nil
}

func GetPriceProduct(data *[]OrderItemsDto, repos ProductRepos) (err error) {
	stock, err := getProducts(*data, repos)
	if err != nil {
		return err
	}

	for i := range *data {
		(*data)[i].Price = stock[i].Price
	}

	return nil
}

func getProducts(items []OrderItemsDto, repos ProductRepos) (stock []product.Product, err error) {
	stock = make([]product.Product, len(items), cap(items))

	for i := range items {
		if stock[i], err = repos.GetProductById(items[i].ProductId); err != nil {
			return nil, err
		}
	}

	return stock, err
}
