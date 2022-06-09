/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:58:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:58:53
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateCarousel 创建轮播图
func CreateCarousel(c *gin.Context) {
	service := service.CarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListCarousels 轮播图列表接口
func ListCarousels(c *gin.Context) {
	service := service.CarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
