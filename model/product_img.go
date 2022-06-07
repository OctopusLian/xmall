/*
 * @Author: neozhang
 * @Date: 2022-06-07 11:42:26
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 11:42:27
 * @Description: 请填写简介
 */
package model

import "github.com/jinzhu/gorm"

// ProductImg 商品图片模型
type ProductImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
