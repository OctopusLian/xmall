/*
 * @Author: neozhang
 * @Date: 2022-06-09 15:51:02
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:51:03
 * @Description: 请填写简介
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateFavorite 创建收藏
func CreateFavorite(c *gin.Context) {
	service := service.FavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowFavorites 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service.FavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteFavorite 删除收藏夹的接口
func DeleteFavorite(c *gin.Context) {
	service := service.FavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
