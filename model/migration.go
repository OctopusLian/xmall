/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-06 22:36:31
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-06 22:36:35
 */
package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}
