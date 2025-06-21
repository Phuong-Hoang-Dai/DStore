package usecase_test

import (
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/user"
	"github.com/Phuong-Hoang-Dai/DStore/user/usecase"
)

func TestJwt(t *testing.T) {
	tests := []struct {
		user    user.User
		wantErr bool
	}{
		{user.User{Name: "Keyone"}, false},
	}

	for _, test := range tests {
		t.Run("Test Jwt", func(t *testing.T) {
			s, err := usecase.GenerateJwt(&test.user)
			if (err != nil) != test.wantErr {
				t.Errorf("Input: %v \n Error : %v", test, err)
			}
			_, err2 := usecase.VerifyJwt(s)
			if (err2 != nil) != test.wantErr {
				t.Errorf("Input: %v \n Error : %v", test, err2)
			}
		})
	}
}
