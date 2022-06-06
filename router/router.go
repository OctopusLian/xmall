/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:53:56
 */
package router

import (
	"xmall/api"
	"xmall/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", api.UserRegister) // 用户注册
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
	}

	return r
}
