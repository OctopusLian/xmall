/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:14:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:15:04
 */
package model

import "github.com/jinzhu/gorm"

// Notice 公告模型 存放公告和邮件模板
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
