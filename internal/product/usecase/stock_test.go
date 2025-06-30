package usecase_test

import (
	"log"
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/internal/product"
	"github.com/Phuong-Hoang-Dai/DStore/internal/product/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetStock(t *testing.T) {
	test := []struct {
		productStock          product.Product
		orderItem             usecase.OrderItemsDto
		quantityStockExpected int
	}{
		{
			productStock:          product.Product{Id: 1, Name: "Milo", Quantity: 10},
			orderItem:             usecase.OrderItemsDto{ProductId: 1, Quantity: 6},
			quantityStockExpected: 4,
		},
		{
			productStock:          product.Product{Id: 2, Name: "Milk", Quantity: 50},
			orderItem:             usecase.OrderItemsDto{ProductId: 2, Quantity: 19},
			quantityStockExpected: 31,
		},
	}

	var mockRepos usecase.MockRepos
	mockRepos.Init()

	log.Println("Create product in mock")
	for _, v := range test {
		mockRepos.CreateProduct(&v.productStock)
	}

	order := []usecase.OrderItemsDto{}
	for _, v := range test {
		order = append(order, v.orderItem)
	}
	log.Println("Call GetStock")

	usecase.GetStock(order, mockRepos)

	for i := range test {
		log.Println("Get product in mock, id: ", i)

		test[i].productStock, _ = mockRepos.GetProductById(test[i].productStock.Id)
		assert.Equal(t, test[i].quantityStockExpected, test[i].productStock.Quantity)
	}
}

func TestRestoreStock(t *testing.T) {
	test := []struct {
		productStock          product.Product
		orderItem             usecase.OrderItemsDto
		quantityStockExpected int
	}{
		{
			productStock:          product.Product{Id: 1, Name: "Milo", Quantity: 10},
			orderItem:             usecase.OrderItemsDto{ProductId: 1, Quantity: 6},
			quantityStockExpected: 16,
		},
		{
			productStock:          product.Product{Id: 2, Name: "Milk", Quantity: 50},
			orderItem:             usecase.OrderItemsDto{ProductId: 2, Quantity: 19},
			quantityStockExpected: 69,
		},
	}

	var mockRepos usecase.MockRepos
	mockRepos.Init()
	log.Println("Create product in mock")
	for _, v := range test {
		mockRepos.CreateProduct(&v.productStock)
	}

	order := []usecase.OrderItemsDto{}
	for _, v := range test {
		order = append(order, v.orderItem)
	}
	log.Println("Call RestoreStock")
	usecase.RestoreStock(order, mockRepos)

	for i := range test {
		log.Println("Get product in mock, id: ", i)

		test[i].productStock, _ = mockRepos.GetProductById(test[i].productStock.Id)
		assert.Equal(t, test[i].quantityStockExpected, test[i].productStock.Quantity)
	}
}
