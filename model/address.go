/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:14:29
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:14:34
 */
package model

import "github.com/jinzhu/gorm"

// Address 收货地址模型
type Address struct {
	gorm.Model
	UserID  uint
	Name    string
	Phone   string
	Address string
}
