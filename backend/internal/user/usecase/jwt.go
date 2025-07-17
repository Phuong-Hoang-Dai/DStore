package usecase

import (
	"time"

	"github.com/Phuong-Hoang-Dai/DStore/configs"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"github.com/Phuong-Hoang-Dai/DStore/utils"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user user.User) (string, error) {
	key := []byte(configs.Cfg.JWTSecret)

	iss := utils.SpaceToDash(configs.Cfg.AppName)
	exp, err := utils.ParseCustomDuration(configs.Cfg.JWTExpireIn)
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.Id,
		"iss":  "Server-" + iss,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(exp).Unix(),
		"role": user.Role,
	})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyJwt(s string) (*jwt.Token, error) {
	key := []byte(configs.Cfg.JWTSecret)

	token, err := jwt.Parse(s, func(t *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func RefreshToken(s string) (string, error) {
	token, err := VerifyJwt(s)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId := int(claims["sub"].(float64))
		u := user.User{Id: userId, Role: claims["role"].(string)}

		if r, err := GenerateJwt(u); err != nil {
			return "", err
		} else {
			return r, nil
		}
	} else {
		return "", user.ErrCannotReadRefreshToken
	}
}
