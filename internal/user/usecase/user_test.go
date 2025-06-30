package usecase_test

import (
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name    string
		uCD     usecase.UserCreateDTO
		uD      usecase.UserDTO
		wantErr bool
	}{
		{
			name:    "test happy case",
			uCD:     usecase.UserCreateDTO{Name: "HoangDai", Password: "Dai12345678"},
			uD:      usecase.UserDTO{Name: "HoangDai", Password: "Dai12345678"},
			wantErr: false,
		},
		{
			name:    "test wrong password",
			uCD:     usecase.UserCreateDTO{Name: "HoangDai", Password: "Dai12345678"},
			uD:      usecase.UserDTO{Name: "HoangDai", Password: "ThisIsWrongPassword"},
			wantErr: true,
		},
	}

	var err error
	var mock usecase.MockRepos
	mock.Init()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.uD.Id, err = usecase.CreateUser(test.uCD, mock)
			assert.NoError(t, err)
			token, err := usecase.Login(test.uD, mock)
			if test.wantErr {
				assert.ErrorIs(t, err, user.ErrUserNameOrPasswordIncorrect)

			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, token, "")
			}
		})
	}
}
