package user_test

import (
	"fmt"
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/user"
	"golang.org/x/crypto/bcrypt"
)

func TestUser(t *testing.T) {
	var user user.WriteUser
	user.Password = []byte("Test12345**")
	user.SetHashPassword()
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, user.Password)

	fmt.Printf("%v \n", err != nil)
	t.Run("Test Process", func(t *testing.T) {

		t.Errorf("HashPassword: %v", user.HashedPassword)
	})
}
