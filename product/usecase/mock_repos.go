package usecase

// import (
// 	"log"

// 	"github.com/Phuong-Hoang-Dai/DStore/product"
// )

// type MockRepos struct{}

// var productList []product.Product

// func (MockRepos) Init() {
// 	productList = []product.Product{}
// }

// func (MockRepos) CreateProduct(data *product.Product) (int, error) {
// 	productList = append(productList, *data)
// 	productList[len(productList)-1].Id = len(productList)
// 	log.Print("MockRepos: Create product ", productList[len(productList)-1])

// 	return productList[len(productList)-1].Id, nil
// }

// func (MockRepos) UpdateProduct(data product.Product) error {
// 	productList[data.Id-1] = data
// 	log.Print("MockRepos: Update product ", data.Id, "/", productList[data.Id-1])

// 	return nil
// }

// func (MockRepos) UpdateProducts(data []product.Product) error {
// 	for i := range data {
// 		productList[data[i].Id-1] = data[i]
// 		log.Print("MockRepos: Update product ", data[i].Id, "/", productList[data[i].Id-1])
// 	}

// 	return nil
// }

// func (MockRepos) GetProductByid(id int, data *product.Product) error {
// 	*data = productList[id-1]
// 	log.Print("MockRepos: Get product ", data.Id, "/", productList[data.Id-1])

// 	return nil
// }

// func (MockRepos) GetProducts(p product.Paging, data *[]product.Product) error {

// 	for i := p.Offset; i < len(productList); i++ {
// 		if i-p.Offset < p.Limit {
// 			*data = append(*data, productList[i])
// 		}
// 	}
// 	log.Print("MockRepos: Get products ", data)

// 	return nil
// }

// func (MockRepos) GetOrderItems(orderId int, data *[]product.OrderItem) error {
// 	return nil
// }

// func (MockRepos) DeleteProduct(id int, data *product.Product) error {
// 	return nil
// }
