package handler

import (
	orderHttp "github.com/Phuong-Hoang-Dai/DStore/order/deliveryhttp"
	productHttp "github.com/Phuong-Hoang-Dai/DStore/product/deliveryhttp"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupHttp(db *gorm.DB) {
	baseUrl := "http://localhost:8080/api/v1"
	productService := productHttp.Init(db)
	orderService := orderHttp.Init(db, baseUrl)

	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			product := v1.Group("/product")
			{
				product.POST("", productService.CreateProduct())
				product.POST("/getstock", productService.GetStock())
				product.POST("/restore", productService.RestoreStock())
				product.POST("/getprice", productService.GetPriceProduct())
				product.GET("/:id", productService.GetProductById())
				product.GET("", productService.GetProducts())
				product.PUT("/:id", productService.UpdateProduct())
				product.DELETE("/:id", productService.DeleteProduct())
			}
			order := v1.Group("/order")
			{
				order.POST("", orderService.CreateOrder())
				order.GET("/:id", orderService.GetOrderById())
				order.GET("", orderService.GetOrders())
				order.DELETE("/:id", orderService.CancelOrder())
				order.PUT("/:id", orderService.UpdateOrderState())
			}
		}
	}
	r.Run()
}
