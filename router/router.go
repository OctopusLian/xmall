/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 17:58:39
 */
package router

import (
	"xmall/api"
	"xmall/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
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
	v2 := r.Group("/api/admin")
	{
		// 管理员注册
		v2.POST("admin/register", api.AdminRegister)
		// 管理员登录
		v2.POST("admin/login", api.AdminLogin)
		//登录验证
		authed2 := v2.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{

		}
	}

	return r
}
