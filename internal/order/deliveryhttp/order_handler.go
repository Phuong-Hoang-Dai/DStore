package deliveryhttp

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Phuong-Hoang-Dai/DStore/internal/order"
	"github.com/Phuong-Hoang-Dai/DStore/internal/order/repos"
	"github.com/Phuong-Hoang-Dai/DStore/internal/order/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepos     usecase.OrderRepository
	productService productHTTPClient
}

func Init(db *gorm.DB, url string, token string) OrderService {
	o := OrderService{
		orderRepos:     repos.NewMysqlOrderRepo(db),
		productService: productHTTPClient{baseURL: fmt.Sprint(url, "/product"), sysToken: token},
	}

	return o
}

func (o OrderService) CreateOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data := []usecase.OrderDTO{}

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		userId := ctx.GetInt("id")
		if userId == 0 {
			responeError(http.StatusBadRequest, order.ErrCannotGetUserId, ctx)
			return
		}
		id, err := usecase.CreateOrder(userId, data, o.orderRepos, o.productService)

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	}
}

func (o OrderService) CancelOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		err = usecase.CancelOrder(id, o.orderRepos, o.productService)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			responeError(http.StatusNotFound, err, ctx)
			return
		}

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, "")
	}
}

func (o OrderService) GetOrderById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		data, err := usecase.GetOrderById(id, o.orderRepos)

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

func (o OrderService) UpdateOrderState() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		data := struct {
			State int `json:"state"`
		}{}

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}

		if err := usecase.UpdateOrder(id, data.State, o.orderRepos); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, "")
	}
}

func (o OrderService) CompleteOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.UpdateOrder(id, order.Completed, o.orderRepos); err != nil {
			if err == gorm.ErrRecordNotFound {
				responeError(http.StatusNotFound, err, ctx)
			} else {
				responeError(http.StatusInternalServerError, err, ctx)
			}
			return
		}

		ctx.JSON(http.StatusOK, "")
	}
}

func (o OrderService) GetOrders() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var p order.Paging
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

		var data []order.Order
		if data, err = usecase.GetOrders(&p, o.orderRepos); err != nil {
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

func (o OrderService) GetHistoryOrders() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var p order.Paging
		var err error

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		userId := ctx.GetInt("id")
		if id != userId {
			responeError(http.StatusBadRequest, order.ErrForbiddened, ctx)
			return
		}

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

		var data []order.Order
		if data, err = usecase.GetHistoryOrders(userId, &p, o.orderRepos); err != nil {
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

func responeError(errCode int, err error, ctx *gin.Context) {
	ctx.JSON(errCode, gin.H{
		"error": err.Error(),
	})
}
