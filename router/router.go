/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 17:22:02
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
		v1.POST("user/login", api.UserLogin)       // 用户登录
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("ping", api.CheckToken) //验证token
			authed.PUT("user", api.UserUpdate) //用户更新操作
		}
	}

	return r
}
