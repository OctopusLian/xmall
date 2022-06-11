/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 19:29:07
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 19:31:54
 */
package service

import (
	"os"
	"strings"
	"time"
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
	"xmall/util"

	"gopkg.in/mail.v2"
)

type EmailService struct {
	UserID   uint   `form:"user_id" json:"user_id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`

	//绑定邮箱
	Token string `form:"token" json:"token"`
}

// Send 发送邮件
func (service *EmailService) Send() serializer.Response {
	code := e.SUCCESS
	var address string
	var notice model.Notice
	token, err := util.GenerateEmailToken(service.UserID, service.OperationType, service.Email, service.Password)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//数据库里 对应邮件id = operation_type+1
	if err := model.DB.First(&notice, service.OperationType+1).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = os.Getenv("VAILD_EMAIL") + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "VaildAddress", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", service.Email)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")抄送
	m.SetHeader("Subject", "CMall")
	m.SetBody("text/html", mailText)

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), 465, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASS"))
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		logging.Info(err)
		code = e.ERROR_SEND_EMAIL
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Vaild 绑定邮箱
func (service *EmailService) Vaild() serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	//验证token
	if service.Token == "" {
		code = e.INVALID_PARAMS
	} else {
		claims, err := util.ParseEmailToken(service.Token)
		if err != nil {
			logging.Info(err)
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if operationType == 1 {
		//1:绑定邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", email).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if operationType == 2 {
		//2：解绑邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", "").Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user, userID).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//3：修改密码
	if operationType == 3 {
		// 加密密码
		if err := user.SetPassword(password); err != nil {
			logging.Info(err)
			code = e.ERROR_FAIL_ENCRYPTION
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Save(&user).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.UPDATE_PASSWORD_SUCCESS
		//返回修改密码成功信息
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
