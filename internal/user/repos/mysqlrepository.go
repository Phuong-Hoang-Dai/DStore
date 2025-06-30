package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/internal/user"
	"github.com/Phuong-Hoang-Dai/DStore/internal/user/usecase"
	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) usecase.UserRepos {
	return mysqlUserRepo{
		DB: db,
	}
}

func (m mysqlUserRepo) CreateUser(data *user.User) (int, error) {
	result := m.DB.Create(&data)

	return data.Id, result.Error
}

func (m mysqlUserRepo) UpdateUser(data user.User) error {
	result := m.DB.Updates(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (m mysqlUserRepo) DeleteUser(id int) error {
	result := m.DB.Delete(&user.User{Id: id})
	if result.RowsAffected == 0 && result.Error == nil {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlUserRepo) GetUserById(id int, data *user.User) error {
	result := m.DB.First(&data, id)

	return result.Error
}

func (m mysqlUserRepo) GetUserByName(name string, data *user.User) error {
	result := m.DB.Where("name = ?", name).First(&data)

	return result.Error
}

func (m mysqlUserRepo) GetUsers(p user.Paging, data *[]user.User) error {
	result := m.DB.Limit(p.Limit).Offset(p.Offset).Find(data)

	return result.Error
}
