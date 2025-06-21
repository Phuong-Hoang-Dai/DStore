package deliveryhttp

import (
	"net/http"
	"strconv"

	"github.com/Phuong-Hoang-Dai/DStore/user"
	"github.com/Phuong-Hoang-Dai/DStore/user/repos"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data user.WriteUser

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}
		user_Repos := repos.NewMysqlUserRepo(db)

		id, err := user_Repos.CreateUser(&data)

		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id product": id,
		})
	}
}

func GetUserById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data user.User

		product_Repos := repos.NewMysqlUserRepo(db)
		err = product_Repos.GetUserById(id, &data)

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

func UpdateUser(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data user.WriteUser

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}

		product_Repos := repos.NewMysqlUserRepo(db)

		if err := product_Repos.UpdateUser(id, &data); err != nil {
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

func DeleteUser(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data user.DeleteUser

		product_Repos := repos.NewMysqlUserRepo(db)

		if err := product_Repos.DeleteUser(id, &data); err != nil {
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

func GetUsers(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var p user.Paging
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

		var data []user.User

		product_Repos := repos.NewMysqlUserRepo(db)

		if err := product_Repos.GetUsers(&p, &data); err != nil {
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
		"message": err.Error(),
	})
}
