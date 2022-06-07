/*
 * @Author: neozhang
 * @Date: 2022-06-07 11:43:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 11:43:07
 * @Description: 请填写简介
 */
package model

import "github.com/jinzhu/gorm"

// ProductParamImg 商品图片模型
type ProductParamImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
