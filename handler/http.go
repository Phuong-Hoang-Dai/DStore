package handler

import (
	"github.com/Phuong-Hoang-Dai/DStore/product/deliveryhttp"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupHttp(db *gorm.DB) {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			product := v1.Group("/product")
			{
				product.POST("", deliveryhttp.CreateProduct(db))
				product.POST("/getstock", deliveryhttp.GetStock(db))
				product.POST("/restore", deliveryhttp.GetStock(db))
				product.GET("/:id", deliveryhttp.GetProductById(db))
				product.GET("", deliveryhttp.GetProducts(db))
				product.PUT("/:id", deliveryhttp.UpdateProduct(db))
				product.DELETE("/:id", deliveryhttp.DeleteProduct(db))
			}
		}
	}
	r.Run()
}
