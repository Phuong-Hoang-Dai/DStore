package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/product"
)

func GetStock(data *[]product.OrderItem, repos ProductRepository) error {
	resetStateItemOrder(data)
	stock := make([]product.Product, len(*data), cap(*data))
	canCreateOrder := true

	for i := range *data {
		stock[i].Id = (*data)[i].ProductId
	}

	err := GetProductsById(&stock, repos)
	if err != nil {
		return err
	}

	for i := range *data {
		if (*data)[i].IsConfirmed = (*data)[i].Quantity <= stock[i].Quantity; !(*data)[i].IsConfirmed {
			canCreateOrder = false
		}
	}

	if !canCreateOrder {
		return product.ErrOutOfStock
	}

	for i := range *data {
		stock[i].Quantity -= (*data)[i].Quantity
		err := repos.UpdateProduct(&stock[i])
		if err != nil {
			return err
		}

		(*data)[i].IsUpdatedStock = true
	}

	return nil
}

func RestoreStock(data *[]product.OrderItem, repos ProductRepository) error {
	resetStateItemOrder(data)
	stock := make([]product.Product, len(*data), cap(*data))

	for i := range *data {
		stock[i].Id = (*data)[i].ProductId
	}

	err := GetProductsById(&stock, repos)
	if err != nil {
		return err
	}

	for i := range *data {
		stock[i].Quantity += (*data)[i].Quantity
		err := repos.UpdateProduct(&stock[i])
		if err != nil {
			return err
		}
		(*data)[i].IsUpdatedStock = true
	}

	return nil
}

func GetProductsById(data *[]product.Product, repos ProductRepository) error {
	for i := range *data {
		err := repos.GetProductByid((*data)[i].Id, &(*data)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func GetOrderTotal(data *[]product.OrderItem, repos ProductRepository) (float32, error) {
	for i := range *data {
		if !(*data)[i].IsConfirmed || !(*data)[i].IsUpdatedStock {
			return 0.0, product.ErrOrderUnvalidated
		}
	}

	stock := make([]product.Product, len(*data), cap(*data))
	for i := range *data {
		stock[i].Id = (*data)[i].ProductId
	}

	err := GetProductsById(&stock, repos)
	if err != nil {
		return 0.0, err
	}

	var total float32
	for i := range *data {
		total += (stock[i].Price)
	}

	return total, nil
}

func resetStateItemOrder(data *[]product.OrderItem) {
	for i := range *data {
		(*data)[i].IsConfirmed = false
		(*data)[i].IsUpdatedStock = false
	}
}
