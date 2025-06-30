package deliveryhttp

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Phuong-Hoang-Dai/DStore/configs"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user/repos"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService struct {
	userRepos usecase.UserRepos
}

func Init(db *gorm.DB) UserService {
	return UserService{
		userRepos: repos.NewMysqlUserRepo(db),
	}
}

func (u UserService) RequireRole(rqRole ...string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": user.ErrMissingRole})
			ctx.Abort()
			return
		}

		userRole := role.(string)
		for _, v := range rqRole {
			if v == userRole {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": user.ErrNotAllowToAccess})
		ctx.Abort()
	}
}

func (u UserService) Authorize() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			responeError(http.StatusBadRequest, user.ErrAuthorizationHeaderMissing, ctx)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			responeError(http.StatusBadRequest, user.ErrAuthorizationHeaderWrongFormat, ctx)
			ctx.Abort()
			return
		}

		tokenStr := parts[1]
		if token, err := usecase.VerifyJwt(tokenStr); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			ctx.Abort()
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userId := int(claims["sub"].(float64))
				ctx.Set("id", userId)
				ctx.Set("role", claims["role"])
			}

			ctx.Next()
		}
	}
}

func (u UserService) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data usecase.UserDTO

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if token, err := usecase.Login(data, u.userRepos); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}

func (u UserService) CreateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data usecase.UserCreateDTO

		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		id, err := usecase.CreateUser(data, u.userRepos)
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id product": id,
		})
	}
}

func (u UserService) GetUserById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		var data usecase.UserResponeDTO
		data, err = usecase.GetUserById(id, u.userRepos)
		if err == gorm.ErrRecordNotFound {
			responeError(http.StatusNotFound, err, ctx)
			return
		}
		if err != nil {
			responeError(http.StatusInternalServerError, err, ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func (u UserService) UpdateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		data := usecase.UserUpdateDTO{}
		if err := ctx.ShouldBind(&data); err != nil {
			responeError(http.StatusBadRequest, err, ctx)
		}

		data.Id = id
		if err := usecase.UpdateUser(data, u.userRepos); err != nil {
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

func (u UserService) DeleteUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responeError(http.StatusBadRequest, err, ctx)
			return
		}

		if err := usecase.DeleteUser(id, u.userRepos); err != nil {
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

func (u UserService) GetUsers() func(ctx *gin.Context) {
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

		var data []usecase.UserResponeDTO

		if data, err = usecase.GetUsers(&p, u.userRepos); err != nil {
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

func (u UserService) SetupJwtSystem() (string, error) {
	sys := usecase.UserCreateDTO{
		Name:     configs.Cfg.SysAccount,
		Password: configs.Cfg.SysPassword,
	}

	var uDAO user.User
	err := u.userRepos.GetUserByName(sys.Name, &uDAO)
	if err == gorm.ErrRecordNotFound {
		id, _ := usecase.CreateUser(sys, u.userRepos)
		uDAO.Id = id
	}

	token, err := usecase.GenerateJwt(uDAO)
	return token, err
}

func responeError(errCode int, err error, ctx *gin.Context) {
	ctx.JSON(errCode, gin.H{
		"error": err.Error(),
	})
}
