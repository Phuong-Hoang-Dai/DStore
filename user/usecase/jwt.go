package usecase

import (
	"log"
	"os"
	"time"

	"github.com/Phuong-Hoang-Dai/DStore/user"
	"github.com/Phuong-Hoang-Dai/DStore/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJwt(user *user.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't Load Env")
		return "", err
	}
	key := []byte(os.Getenv("JWT_SECRET"))

	iss := utils.SpaceToDash(os.Getenv("APP_NAME"))
	exp, err := utils.ParseCustomDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.Name,
		"iss":  "Server-" + iss,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(exp).Unix(),
		"role": user.RoleId,
	})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyJwt(s string) (jwt.Claims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
