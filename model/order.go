/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:16:57
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:17:01
 */
package model

import "github.com/jinzhu/gorm"

// Order 订单模型
type Order struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	Num          uint
	OrderNum     uint64
	AddressName  string
	AddressPhone string
	Address      string
	Type         uint
}
