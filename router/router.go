/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:30:27
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
		v1.Use(middleware.JWT())
		{
			v1.GET("/ping", api.CheckToken)             //验证token
			v1.PUT("/user", api.UserUpdate)             //用户更新操作
			v1.GET("/notices", api.ShowNotice)          //查看公告详情
			v1.GET("/products", api.ListProducts)       //查询商品列表
			v1.GET("/products/:id", api.ShowProduct)    //查询一个商品的详情
			v1.POST("/searches", api.SearchProducts)    //搜索商品
			v1.POST("/addresses", api.CreateAddress)    //创建收获地址
			v1.GET("/addresses/:id", api.ShowAddresses) //展示收货地址
			v1.PUT("/addresses", api.UpdateAddress)     //修改收货地址
			v1.DELETE("/addresses", api.DeleteAddress)  //删除收货地址
		}

		v0 := r.Group("/api/v0")
		{
			v0.POST("/register", api.AdminRegister) // 管理员注册
			v0.POST("/login", api.AdminLogin)       // 管理员登录
			//登录验证
			v0.Use(middleware.JWTAdmin())
			{
				v0.POST("/notices", api.CreateNotice) //创建公告
				v0.PUT("/notices", api.UpdateNotice)  //更新公告
			}
		}
	}

	return r
}
