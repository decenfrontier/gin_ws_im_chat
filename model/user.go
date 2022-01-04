package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Email          string //`gorm:"unique"`
	Avatar         string `gorm:"size:1000"`
	Phone          string
	Status         string
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

//SetPassword 设置密码
func (this *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	this.PasswordDigest = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (this *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
	return err == nil
}
