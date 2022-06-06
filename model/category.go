/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:17:32
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:17:37
 */
package model

import "github.com/jinzhu/gorm"

// Category 分类模型
type Category struct {
	gorm.Model
	CategoryID   uint
	CategoryName string
	Num          uint
}
