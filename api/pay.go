/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 19:22:57
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 19:23:00
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// InitPay 初始化支付
func InitPay(c *gin.Context) {
	service := service.PayService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ConfirmPay 接收FM支付回调接口
func ConfirmPay(c *gin.Context) {
	service := service.PayService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Confirm()
		c.String(200, "success")
	} else {
		c.String(200, "success")
		logging.Info(err)
	}
}
