package usecase

import (
	"errors"

	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
)

type MockRepos struct{}

var userList []user.User

func (MockRepos) Init() {
	userList = []user.User{}
}

func (MockRepos) CreateUser(data *user.User) (int, error) {
	userList = append(userList, *data)
	userList[len(userList)-1].Id = len(userList)

	return userList[len(userList)-1].Id, nil
}

func (MockRepos) UpdateUser(data user.User) error {
	userList[data.Id-1] = data

	return nil
}

func (MockRepos) GetUserById(id int, data *user.User) error {
	*data = userList[id-1]
	return nil
}

func (MockRepos) GetUserByName(name string, data *user.User) error {
	for _, v := range userList {
		if v.Name == name {
			*data = v
			return nil
		}
	}
	return errors.New("can't find user")
}

func (MockRepos) GetUsers(p user.Paging, data *[]user.User) error {
	for i := p.Offset; i < len(userList); i++ {
		if i-p.Offset < p.Limit {
			*data = append(*data, userList[i])
		}
	}

	return nil
}

func (MockRepos) DeleteUser(id int) error {
	return nil
}
