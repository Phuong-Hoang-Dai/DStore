package deliveryhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Phuong-Hoang-Dai/DStore/order"
	"github.com/Phuong-Hoang-Dai/DStore/order/repos"
	"github.com/Phuong-Hoang-Dai/DStore/order/usecase"
	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data := []order.OrderItem{}

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		result, err := makeRequest("POST", "http://localhost:8080/product/getstock", data)
		data, ok := result.([]order.OrderItem)
		if !ok {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		// jsonData, err := json.Marshal(data)
		// if err != nil {
		// 	responeError(http.StatusBadRequest, err, ctx)
		// }

		// req, err := http.NewRequest("POST", "http://localhost:8080/product/getstock", bytes.NewBuffer(jsonData))
		// if err != nil {
		// 	responeError(http.StatusInternalServerError, err, ctx)
		// }
		// req.Header.Set("Content-Type", "application/json")

		// client := &http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	responeError(http.StatusInternalServerError, err, ctx)
		// }
		// defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	responeError(http.StatusInternalServerError, err, ctx)

		// }

		// var raw map[string]json.RawMessage
		// if err := json.Unmarshal(body, &raw); err != nil {
		// 	responeError(http.StatusInternalServerError, err, ctx)
		// }

		// errMsg := ""
		// if rawErr, ok := raw["error"]; ok {
		// 	json.Unmarshal(rawErr, &errMsg)
		// }

		// if rawErr, ok := raw["data"]; ok {
		// 	json.Unmarshal(rawErr, &data)
		// }

		// orderRepos := repos.NewMysqlOrderRepo(db)
		// if errMsg != "" {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": errMsg,
		// 		"data":  data,
		// 	})
		// 	return
		// }

		orderRepos := repos.NewMysqlOrderRepo(db)
		dataOrder := order.Order{Items: data}
		id, err := usecase.CreateOrder(&dataOrder, orderRepos)

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": "",
			"id":    id,
		})
	}
}

func GetOrderById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data order.Order

		orderRepos := repos.NewMysqlOrderRepo(db)
		err = orderRepos.GetOrderById(id, &data)

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

func UpdateOrderState(db *gorm.DB) func(ctx *gin.Context) {
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

		var order order.Order
		order.State = data.State
		order.Id = id

		orderRepos := repos.NewMysqlOrderRepo(db)

		if err := orderRepos.UpdateOrder(&order); err != nil {
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

func CancelOrder(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data order.Order
		data.Id = id

		orderRepos := repos.NewMysqlOrderRepo(db)
		orderRepos.GetOrderById(data.Id, &data)

		itemsData := data.Items
		jsonData, err := json.Marshal(itemsData)
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/product/restore", bytes.NewBuffer(jsonData))
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)

		}

		var raw map[string]json.RawMessage
		if err := json.Unmarshal(body, &raw); err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
		}

		errMsg := ""
		if rawErr, ok := raw["error"]; ok {
			json.Unmarshal(rawErr, &errMsg)
		}

		if errMsg != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": errMsg,
			})
			return
		}

		data.State = order.Cancelled
		orderRepos.UpdateOrder(&data)

		if err := orderRepos.DeleteOrder(id, &data); err != nil {
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

func GetOrders(db *gorm.DB) func(ctx *gin.Context) {
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

		var data []order.Order

		orderRepos := repos.NewMysqlOrderRepo(db)

		if err := orderRepos.GetOrders(&data); err != nil {
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

func makeRequest(method string, url string, data any) (any, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err

	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	errMsg := ""
	if rawErr, ok := raw["error"]; ok {
		json.Unmarshal(rawErr, &errMsg)
	}

	if rawErr, ok := raw["data"]; ok {
		json.Unmarshal(rawErr, &data)
	}

	return data, errors.New(errMsg)
}
