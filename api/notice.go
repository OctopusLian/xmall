/*
 * @Author: neozhang
 * @Date: 2022-06-07 18:06:00
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 18:06:01
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

//ShowNotice 公告详情
func ShowNotice(c *gin.Context) {
	service := service.NoticeService{}
	res := service.Show()
	c.JSON(200, res)
}

//CreateNotice 创建公告
func CreateNotice(c *gin.Context) {
	service := service.NoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UpdateNotice 更新公告
func UpdateNotice(c *gin.Context) {
	service := service.NoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
