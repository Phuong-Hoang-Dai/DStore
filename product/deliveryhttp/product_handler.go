package deliveryhttp

import (
	"net/http"
	"strconv"

	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/Phuong-Hoang-Dai/DStore/product/repos"
	"github.com/Phuong-Hoang-Dai/DStore/product/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data product.Product

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}
		productRepos := repos.NewMysqlProductRepo(db)

		id, err := productRepos.CreateProduct(&data)

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id product": id,
		})
	}
}

func GetProductById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data product.Product

		productRepos := repos.NewMysqlProductRepo(db)
		err = productRepos.GetProductByid(id, &data)

		if err == gorm.ErrRecordNotFound {
			responeError(http.StatusNotFound, err, ctx)
			return
		}
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, data)
	}
}

func GetProductsById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data []product.Product
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		productRepos := repos.NewMysqlProductRepo(db)
		err := usecase.GetProductsById(&data, productRepos)

		if err == gorm.ErrRecordNotFound {
			responeError(http.StatusNotFound, err, ctx)
			return
		}

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, data)
	}
}

func UpdateProduct(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data product.Product
		data.Id = id

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}

		productRepos := repos.NewMysqlProductRepo(db)

		if err := productRepos.UpdateProduct(&data); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, data)
	}
}

func DeleteProduct(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data product.Product

		productRepos := repos.NewMysqlProductRepo(db)

		if err := productRepos.DeleteProduct(id, &data); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, id)
	}
}

func GetProducts(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var p product.Paging
		var err error
		p.Limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "0"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}
		p.Offset, err = strconv.Atoi(ctx.DefaultQuery("offset", "0"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data []product.Product

		productRepos := repos.NewMysqlProductRepo(db)

		if err := productRepos.GetProducts(&p, &data); err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"limit":   p.Limit,
			"offfset": p.Offset,
			"data":    data,
		})
	}
}

func GetStock(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data := []product.OrderItem{}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}

		productRepos := repos.NewMysqlProductRepo(db)
		err := usecase.GetStock(&data, productRepos)

		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
				"data":  data,
			})
			return
		}

		if err == product.ErrOutOfStock {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": product.ErrOutOfStock,
				"data":  data,
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": "",
			"data":  data,
		})
	}
}

func RestoreStock(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data := []product.OrderItem{}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}

		productRepos := repos.NewMysqlProductRepo(db)
		err := usecase.RestoreStock(&data, productRepos)

		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
				"data":  data,
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": "",
		})
	}
}

func GetOrderTotal(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data := []product.OrderItem{}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}
		productRepos := repos.NewMysqlProductRepo(db)
		total, err := usecase.GetOrderTotal(&data, productRepos)

		if err == product.ErrOrderUnvalidated {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"data":  0,
			})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"data":  "",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "",
			"data":  total,
		})
	}
}

func responeError(errCode int, err error, ctx *gin.Context) {
	ctx.JSON(errCode, gin.H{
		"error": err.Error(),
	})
}
