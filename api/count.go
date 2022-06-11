/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 19:04:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 19:04:48
 */
package api

import (
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// ShowCount 获取数量服务
func ShowCount(c *gin.Context) {
	service := service.ShowCountService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
