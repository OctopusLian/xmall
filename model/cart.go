/*
 * @Author: neozhang
 * @Date: 2022-06-07 11:37:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:34:52
 * @Description: 请填写简介
 */
package model

import "github.com/jinzhu/gorm"

// Cart 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Num       uint
	MaxNum    uint
	Check     bool
}
