/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 06:51:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 23:50:10
 */
package main

import (
	"xmall/conf"
	"xmall/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	//初始化路由
	r := router.NewRouter()
	r.Run(":3000")
}
