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

type ProductService struct {
	productRepos usecase.ProductRepos
}

func Init(db *gorm.DB) ProductService {
	return ProductService{productRepos: repos.NewMysqlProductRepo(db)}
}

func (p ProductService) CreateProduct() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data product.Product

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		id, err := usecase.CreateProduct(data, p.productRepos)
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	}
}

func (p ProductService) GetProductById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data product.Product
		if data, err = usecase.GetProductById(id, p.productRepos); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
				return
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
				return
			}
		}

		ctx.JSON(http.StatusOK, data)
	}
}

func (p ProductService) UpdateProduct() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		data := product.Product{}
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}
		data.Id = id

		if err := usecase.UpdateProduct(data, p.productRepos); err != nil {
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

func (p ProductService) DeleteProduct() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.DeleteProduct(id, p.productRepos); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	}
}

func (p ProductService) GetProducts() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging product.Paging
		var err error

		paging.Limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "0"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}
		paging.Offset, err = strconv.Atoi(ctx.DefaultQuery("offset", "0"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data []product.Product
		if data, err = usecase.GetProducts(&paging, p.productRepos); err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    data,
			"limit":   paging.Limit,
			"offfset": paging.Offset,
		})
	}
}

func (p ProductService) GetStock() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data []usecase.OrderItemsDto
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.GetStock(data, p.productRepos); err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				responeError(http.StatusNotFound, err, ctx)
			case product.ErrOutOfStock:
				responeError(http.StatusBadRequest, err, ctx)
			default:
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, "")
	}
}

func (p ProductService) RestoreStock() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data []usecase.OrderItemsDto
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.RestoreStock(data, p.productRepos); err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				responeError(http.StatusNotFound, err, ctx)
			case product.ErrOutOfStock:
				responeError(http.StatusBadRequest, err, ctx)
			default:
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": "",
		})
	}
}

func (p ProductService) GetPriceProduct() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data []usecase.OrderItemsDto
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.GetPriceProduct(&data, p.productRepos); err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				responeError(http.StatusNotFound, err, ctx)
			case product.ErrOutOfStock:
				responeError(http.StatusBadRequest, err, ctx)
			default:
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func responeError(errCode int, err error, ctx *gin.Context) {
	ctx.JSON(errCode, gin.H{
		"error": err.Error(),
	})
}
