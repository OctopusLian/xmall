/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:11:22
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:11:22
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateCart 加入购物车
func CreateCart(c *gin.Context) {
	service := service.CartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowCarts 购物车详情接口
func ShowCarts(c *gin.Context) {
	service := service.CartService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateCart 修改购物车信息
func UpdateCart(c *gin.Context) {
	service := service.CartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	service := service.CartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
