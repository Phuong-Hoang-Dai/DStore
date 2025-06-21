package user

import "time"

type Role struct {
	Id        int       `json:"roleId" gorm:"column:roleId"`
	Name      string    `json:"roleName" gorm:"column:name"`
	Create_at time.Time `json:"create_at" gorm:"column:create_at"`
	Update_at time.Time `json:"update_at" gorm:"column:update_at"`
	Delete_at time.Time `json:"delete_at" gorm:"column:delete_at"`
}

func (Role) GetTableName() string { return "roles" }
