/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:39:37
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:50:39
 */
package api

import (
	"encoding/json"
	"fmt"
	"xmall/logging"
	"xmall/serializer"
	"xmall/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.Response{
				Status: 40001,
				Msg:    "",
				Error:  fmt.Sprint(err),
			}
		}
	}
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
	session := sessions.Default(c)
	userID := session.Get("userId")
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register(userID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
