package usecase

import (
	"github.com/Phuong-Hoang-Dai/DStore/user"
)

type UserRepository interface {
	Login(data *user.User) (string, error)
	CreateUser(data *user.WriteUser) (int, error)
	UpdateUser(id int, data *user.WriteUser) error
	GetUserById(id int, data *user.User) error
	GetUsers(p *user.Paging, data *[]user.User) error
	DeleteUser(id int, data *user.DeleteUser) error
}
