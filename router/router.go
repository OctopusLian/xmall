/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-08 22:25:12
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
		v1.POST("/user/register", api.UserRegister) // 用户注册
		v1.POST("/user/login", api.UserLogin)       // 用户登录

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("/ping", api.CheckToken)     //验证token
			authed.PUT("/user", api.UserUpdate)     //用户更新操作
			v1.GET("/notices", api.ShowNotice)      //查看公告详情
			v1.GET("products", api.ListProducts)    //查询商品列表
			v1.GET("products/:id", api.ShowProduct) //查询一个商品的详情
			v1.POST("searches", api.SearchProducts) //搜索商品
		}

		v1.POST("/admin/register", api.AdminRegister) // 管理员注册
		v1.POST("/admin/login", api.AdminLogin) // 管理员登录
		//登录验证
		authed2 := v1.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{
			authed2.POST("/notices", api.CreateNotice) //创建公告
			authed2.PUT("/notices", api.UpdateNotice)  //更新公告
		}
	}

	return r
}
