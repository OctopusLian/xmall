/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:16:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:16:08
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string
	CategoryID    int
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
}
