package service

import (
	"chat/model"
	"chat/serializer"
)

type UserRegisterService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"user_name" form:"user_name"`
}

func (this *UserRegisterService) Register() serializer.Response {
	var user model.User
	count := 0
	model.DB.Model(&model.User{}).Where("user_name=?", this.UserName).First(&user).Count(&count)
	if count != 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已存在",
		}
	}
	user = model.User{
		UserName: this.UserName,
	}
	err := user.SetPassword(this.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "加密出错",
		}
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
