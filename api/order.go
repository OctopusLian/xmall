/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:29:32
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:29:33
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	service := service.OrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListOrder 订单列表接口
func ListOrders(c *gin.Context) {
	service := service.OrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowOrder 订单详情接口
func ShowOrder(c *gin.Context) {
	service := service.OrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("num"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
