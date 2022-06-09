/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:38:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 22:33:47
 */
package conf

import (
	"xmall/cache"
	"xmall/model"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	model.Database("root:mysql123@tcp(127.0.0.1:3306)/xmall?charset=utf8&parseTime=True&loc=Local")
	cache.Redis()
}
