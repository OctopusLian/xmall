/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 06:51:57
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-08 22:00:51
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
	r.Run(":8001")
}
