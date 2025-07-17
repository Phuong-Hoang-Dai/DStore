package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
)

type UserRepos interface {
	CreateUser(data *user.User) (int, error)
	UpdateUser(data user.User) error
	GetUserById(id int, data *user.User) error
	GetUserByName(name string, data *user.User) error
	GetUsers(p user.Paging, data *[]user.User) error
	DeleteUser(id int) error
}
