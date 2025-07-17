package handler

import (
	"log"
	"time"

	orderHttp "github.com/Phuong-Hoang-Dai/DStore/internal/order/deliveryhttp"
	productHttp "github.com/Phuong-Hoang-Dai/DStore/internal/product/deliveryhttp"
	userHttp "github.com/Phuong-Hoang-Dai/DStore/internal/user/deliveryhttp"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupHttp(db *gorm.DB) {
	baseUrl := "http://localhost:8069/api/v1"
	userService := userHttp.Init(db)
	sysToken, err := userService.SetupJwtSystem()

	if err != nil {
		log.Fatal(err)
	}
	productService := productHttp.Init(db)
	orderService := orderHttp.Init(db, baseUrl, sysToken)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("/login", userService.Login())
				user.POST("", userService.CreateUser())
				user.PUT("/:id", userService.Authorize(), userService.RequireRole("user", "system"), userService.UpdateUser())
				user.GET("/:id", userService.GetUserById())
				user.GET("", userService.GetUsers())
				user.DELETE("/:id", userService.Authorize(), userService.RequireRole("user", "system"), userService.DeleteUser())
			}

			product := v1.Group("/product")
			{
				product.POST("", userService.Authorize(), userService.RequireRole("admin", "system"), productService.CreateProduct())
				product.POST("/getstock", userService.Authorize(), userService.RequireRole("system"), productService.GetStock())
				product.POST("/restore", userService.Authorize(), userService.RequireRole("system"), productService.RestoreStock())
				product.GET("", productService.GetProducts())
				product.POST("/getprice", productService.GetPriceProduct())
				product.GET("/:id", productService.GetProductById())
				product.PUT("/:id", userService.Authorize(), userService.RequireRole("admin", "system"), productService.UpdateProduct())
				product.DELETE("/:id", userService.Authorize(), userService.RequireRole("admin", "system"), productService.DeleteProduct())
			}
			order := v1.Group("/order", userService.Authorize())
			{
				order.POST("", userService.RequireRole("user", "system"), orderService.CreateOrder())
				order.GET("/:id", userService.RequireRole("user", "system"), orderService.GetOrderById())
				order.GET("", userService.RequireRole("system", "admin"), orderService.GetOrders())
				order.GET("/history/:id", userService.RequireRole("user"), orderService.GetHistoryOrders())
				order.DELETE("/:id", userService.RequireRole("user"), orderService.CancelOrder())
				order.PUT("/:id", userService.RequireRole("admin", "system"), orderService.UpdateOrderState())
				order.POST("/completeOrder/:id", userService.RequireRole("admin", "system"), orderService.CompleteOrder())
			}
		}
	}
	r.Run(":8069")
}
