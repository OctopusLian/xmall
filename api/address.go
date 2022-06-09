/*
 * @Author: neozhang
 * @Date: 2022-06-09 15:17:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:19:01
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateAddress 新建收货地址
func CreateAddress(c *gin.Context) {
	service := service.AddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowAddresses 展示收货地址
func ShowAddresses(c *gin.Context) {
	service := service.AddressService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	service := service.AddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteAddress 删除收货地址
func DeleteAddress(c *gin.Context) {
	service := service.AddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
