/*
 * @Author: neozhang
 * @Date: 2022-06-07 17:51:53
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 17:53:33
 * @Description: 请填写简介
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
	"xmall/util"

	"github.com/jinzhu/gorm"
)

type AdminService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Login 管理员登录函数
func (service *AdminService) Login() serializer.Response {
	var admin model.Admin
	code := e.SUCCESS

	if err := model.DB.Where("user_name = ?", service.UserName).First(&admin).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_USER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.ERROR_NOT_EXIST_USER
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if admin.CheckPassword(service.Password) == false {
		code = e.ERROR_NOT_COMPARE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(service.UserName, service.Password, 1)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildAdmin(admin), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证表单
func (service *AdminService) Valid() *serializer.Response {
	var code int
	if service.PasswordConfirm != service.Password {
		code = e.ERROR_NOT_COMPARE_PASSWORD
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	count := 0
	model.DB.Model(&model.Admin{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		code = e.ERROR_EXIST_USER
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return nil
}

// Register 管理员注册
func (service *AdminService) Register() *serializer.Response {
	admin := model.Admin{
		UserName: service.UserName,
	}
	code := e.SUCCESS
	// 表单验证
	if res := service.Valid(); res != nil {
		return res
	}

	// 加密密码
	if err := admin.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	if err := model.DB.Create(&admin).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
