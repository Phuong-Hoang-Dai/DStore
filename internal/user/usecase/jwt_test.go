package usecase_test

import (
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	tests := []struct {
		user user.User
	}{
		{user.User{Id: 999666, Name: "Keyone"}},
	}

	for _, test := range tests {
		t.Run("Test Jwt", func(t *testing.T) {
			s, err := usecase.GenerateJwt(test.user)
			assert.NoError(t, err)
			token, err := usecase.VerifyJwt(s)
			assert.NoError(t, err)

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userId := int(claims["sub"].(float64))
				assert.Equal(t, test.user.Id, userId)
			}

		})
	}
}
