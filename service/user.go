/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:54
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:45:26
 */
package service

import (
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
	"xmall/status"
)

type UserService struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Challenge string `form:"challenge" json:"challenge"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

// Register 用户注册
func (service *UserService) Register(userID interface{}) *serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}
	code := status.SUCCESS

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = status.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}

	user.Avatar = "img/avatar/avatar1.jpg"

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = status.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    status.GetMsg(code),
	}
}
