/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:39:37
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 19:55:59
 */
package api

import (
	"encoding/json"
	"fmt"
	"xmall/logging"
	"xmall/serializer"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("userId")
	// logging.Info("userID:", userID)
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	// session := sessions.Default(c)
	// userID := session.Get("userId")
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UserUpdate 用户修改信息
func UserUpdate(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// CheckToken
func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

// SendEmail 发送邮件接口
func SendEmail(c *gin.Context) {
	var service service.EmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Send()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// VaildEmail 绑定和解绑邮箱接口
func VaildEmail(c *gin.Context) {
	var service service.EmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
