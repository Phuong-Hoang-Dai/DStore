package user_test

import (
	"fmt"
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"golang.org/x/crypto/bcrypt"
)

func TestUser(t *testing.T) {
	var user user.User
	Password := []byte("Test12345**")
	user.SetHashPassword(Password)
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, Password)

	fmt.Printf("%v \n", err != nil)
	t.Run("Test Process", func(t *testing.T) {

		t.Errorf("HashPassword: %v", user.HashedPassword)
	})
}
