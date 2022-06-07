/*
 * @Author: neozhang
 * @Date: 2022-06-07 11:41:44
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 11:41:44
 * @Description: 请填写简介
 */
package model

import "github.com/jinzhu/gorm"

// Favorite 收藏夹模型
type Favorite struct {
	gorm.Model
	UserID    uint
	ProductID uint
}
