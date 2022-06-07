/*
 * @Author: neozhang
 * @Date: 2022-06-07 11:42:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 11:42:46
 * @Description: 请填写简介
 */
package model

import "github.com/jinzhu/gorm"

// ProductInfoImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
