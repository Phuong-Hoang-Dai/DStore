package repos

import (
	"github.com/Phuong-Hoang-Dai/DStore/user"
	"github.com/Phuong-Hoang-Dai/DStore/user/usecase"
	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) usecase.UserRepository {
	return mysqlUserRepo{
		DB: db,
	}
}

func (m mysqlUserRepo) Login(data *user.User) (string, error) {
	result := m.DB.Table(user.UserTableName).First(&data, "name = ?", data.Name)
	if result.Error != nil {
		return "", result.Error
	}

	err := data.VerifyPassword()
	if err != nil {
		return "", err
	}

	token, err := usecase.GenerateJwt(data)

	return token, err
}

func (m mysqlUserRepo) CreateUser(data *user.WriteUser) (int, error) {
	err := data.Validate()
	if err != nil {
		return -1, err
	}
	data.SetHashPassword()
	result := m.DB.Table(user.UserTableName).Create(&data)

	return data.Id, result.Error
}

func (m mysqlUserRepo) UpdateUser(id int, data *user.WriteUser) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	result := m.DB.Table(user.UserTableName).Updates(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}

	return result.Error
}

func (m mysqlUserRepo) DeleteUser(id int, data *user.DeleteUser) error {
	result := m.DB.Table(user.UserTableName).Where("id = ?", id).Delete(&data)
	if result.RowsAffected == 0 && result.Error == nil {
		result.Error = gorm.ErrRecordNotFound
	}
	return result.Error
}

func (m mysqlUserRepo) GetUserById(id int, data *user.User) error {
	result := m.DB.Table(user.UserTableName).First(&data, id)

	return result.Error
}

func (m mysqlUserRepo) GetUsers(p *user.Paging, data *[]user.User) error {
	p.Process()
	result := m.DB.Table(user.UserTableName).Limit(p.Limit).Offset(p.Offset).Find(data)

	return result.Error
}
