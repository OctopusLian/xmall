/*
 * @Author: neozhang
 * @Date: 2022-06-07 17:56:02
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 17:56:03
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// AdminRegister 管理员登录接口
func AdminRegister(c *gin.Context) {
	var service service.AdminService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// AdminLogin 管理员登录接口
func AdminLogin(c *gin.Context) {
	var service service.AdminService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
