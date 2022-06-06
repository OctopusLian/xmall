/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:38:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-06 22:38:22
 */
package conf

import (
	"os"

	"xmall/model"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}
