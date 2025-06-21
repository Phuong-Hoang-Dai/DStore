package user

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id              int       `json:"userId" gorm:"column:userId"`
	Name            string    `json:"name" gorm:"column:name"`
	Email           string    `json:"email" gorm:"column:email"`
	Password        []byte    `json:"password"`
	HashedPassword  []byte    `gorm:"column:password"`
	Create_at       time.Time `json:"create_at" gorm:"column:create_at"`
	Update_at       time.Time `json:"update_at" gorm:"column:update_at"`
	Delete_at       time.Time `json:"delete_at" gorm:"column:delete_at"`
	RoleId          int       `json:"roleId" gorm:"column:roleId"`
	RoleGrantedById int       `json:"roleGrantedById" gorm:"column:roleId"`
}

type WriteUser struct {
	Id              int    `json:"userId" gorm:"column:userId"`
	Name            string `json:"name" gorm:"column:name"`
	Email           string `json:"email" gorm:"column:email"`
	Password        []byte `json:"password"`
	HashedPassword  []byte `gorm:"column:password"`
	RoleId          int    `json:"roleId" gorm:"column:roleId"`
	RoleGrantedById int    `json:"roleGrantedById" gorm:"column:roleId"`
}

type DeleteUser struct {
	Id        int       `json:"userId" gorm:"column:userId"`
	Delete_at time.Time `json:"delete_at" gorm:"column:delete_at"`
}

func (User) GetTableName() string      { return "users" }
func (WriteUser) GetTableName() string { return "users" }

func (user User) VerifyPassword() error {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, user.Password)
	return err
}

func (user *WriteUser) SetHashPassword() {
	hashedPassword := hashPassword(user.Password)
	user.HashedPassword = hashedPassword
}

func (user WriteUser) Validate() error {
	return nil
}

func hashPassword(pw []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("hashPassword error: %v", err)
	}
	return hashedPassword
}

type Paging struct {
	Limit  int
	Offset int
}

func (p *Paging) Process() {
	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
	if p.Limit < 0 {
		p.Limit = 0
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
}
