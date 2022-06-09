/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:44:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:30:44
 */
package router

import (
	"xmall/api"
	"xmall/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//用户
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
			v1.GET("favorites/:id", api.ShowFavorites)  //新建收藏
			v1.POST("favorites", api.CreateFavorite)    //创建收藏
			v1.DELETE("favorites", api.DeleteFavorite)  //删除收藏
			v1.POST("carts", api.CreateCart)            //加入购物车
			v1.GET("carts/:id", api.ShowCarts)          //购物车详情
			v1.PUT("carts", api.UpdateCart)             //修改购物车信息
			v1.DELETE("carts", api.DeleteCart)          //删除购物车
			v1.POST("orders", api.CreateOrder)          //创建订单
			v1.GET("user/:id/orders", api.ListOrders)   //订单列表
			v1.GET("orders/:num", api.ShowOrder)        //订单详情
		}
	}
	//管理员
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

	return r
}
