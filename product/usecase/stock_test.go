package usecase_test

// import (
// 	"log"
// 	"testing"

// 	"github.com/Phuong-Hoang-Dai/DStore/product"
// 	"github.com/Phuong-Hoang-Dai/DStore/product/usecase"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetStock(t *testing.T) {
// 	test := []struct {
// 		productStock          product.Product
// 		orderItem             product.OrderItem
// 		quantityStockExpected int
// 	}{
// 		{
// 			productStock:          product.Product{Id: 1, Name: "Milo", Quantity: 10},
// 			orderItem:             product.OrderItem{ProductId: 1, Quantity: 6, OrderId: 1},
// 			quantityStockExpected: 4,
// 		},
// 		{
// 			productStock:          product.Product{Id: 2, Name: "Milk", Quantity: 50},
// 			orderItem:             product.OrderItem{ProductId: 2, Quantity: 19, OrderId: 1},
// 			quantityStockExpected: 31,
// 		},
// 	}

// 	var mockRepos usecase.MockRepos
// 	mockRepos.Init()

// 	log.Println("Create product in mock")
// 	for _, v := range test {
// 		mockRepos.CreateProduct(&v.productStock)
// 	}

// 	order := []product.OrderItem{}
// 	for _, v := range test {
// 		order = append(order, v.orderItem)
// 	}
// 	log.Println("Call GetStock")

// 	usecase.GetStock(order[0].Order.Id, mockRepos)

// 	for i := range test {
// 		log.Println("Get product in mock, id: ", i)

// 		mockRepos.GetProductByid(test[i].productStock.Id, &test[i].productStock)
// 		assert.Equal(t, test[i].quantityStockExpected, test[i].productStock.Quantity)
// 	}
// }

// func TestRestoreStock(t *testing.T) {
// 	test := []struct {
// 		productStock          product.Product
// 		orderItem             product.OrderItem
// 		quantityStockExpected int
// 	}{
// 		{
// 			productStock:          product.Product{Id: 1, Name: "Milo", Quantity: 10},
// 			orderItem:             product.OrderItem{ProductId: 1, Quantity: 6},
// 			quantityStockExpected: 16,
// 		},
// 		{
// 			productStock:          product.Product{Id: 2, Name: "Milk", Quantity: 50},
// 			orderItem:             product.OrderItem{ProductId: 2, Quantity: 19},
// 			quantityStockExpected: 69,
// 		},
// 	}

// 	var mockRepos usecase.MockRepos
// 	mockRepos.Init()
// 	log.Println("Create product in mock")
// 	for _, v := range test {
// 		mockRepos.CreateProduct(&v.productStock)
// 	}

// 	order := []product.OrderItem{}
// 	for _, v := range test {
// 		order = append(order, v.orderItem)
// 	}
// 	log.Println("Call RestoreStock")

// 	for i := range test {
// 		log.Println("Get product in mock, id: ", i)

// 		mockRepos.GetProductByid(test[i].productStock.Id, &test[i].productStock)
// 		assert.Equal(t, test[i].quantityStockExpected, test[i].productStock.Quantity)
// 	}
// }

// func TestOrderTotal(t *testing.T) {
// 	test := []struct {
// 		productStock  product.Product
// 		orderItem     product.OrderItem
// 		totalExpected float32
// 	}{
// 		{
// 			productStock:  product.Product{Id: 1, Name: "Milo", Quantity: 10, Price: 10},
// 			orderItem:     product.OrderItem{ProductId: 1, Quantity: 6},
// 			totalExpected: 60,
// 		},
// 		{
// 			productStock:  product.Product{Id: 2, Name: "Milk", Quantity: 50, Price: 5},
// 			orderItem:     product.OrderItem{ProductId: 2, Quantity: 19},
// 			totalExpected: 95,
// 		},
// 	}
// 	var expectedTotal float32
// 	for _, v := range test {
// 		expectedTotal += v.totalExpected
// 	}

// 	var mockRepos usecase.MockRepos
// 	mockRepos.Init()
// 	log.Println("Create product in mock")
// 	for _, v := range test {
// 		mockRepos.CreateProduct(&v.productStock)
// 	}

// 	order := []product.OrderItem{}
// 	for _, v := range test {
// 		order = append(order, v.orderItem)
// 	}
// 	log.Println("Call RestoreStock")

// 	actualTotal, err := usecase.GetPriceProduct(1, mockRepos)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedTotal, actualTotal)

// }
