/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:07:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:10:22
 * @Description: 请填写简介
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type CartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

// Create 创建购物车
func (service *CartService) Create() serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, service.ProductID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if product == (model.Product{}) {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var cart model.Cart
	model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart)
	//如果不存在该购物车则创建
	if cart == (model.Cart{}) {
		cart = model.Cart{
			UserID:    service.UserID,
			ProductID: service.ProductID,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}

		err = model.DB.Create(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildCart(cart, product),
		}
	} else if cart.Num < cart.MaxNum { //如果存在该购物车且num小于maxnum
		cart.Num++
		err = model.DB.Save(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: 201,
			Msg:    "商品已在购物车，数量+1",
			Data:   serializer.BuildCart(cart, product),
		}
	} else {
		return serializer.Response{
			Status: 202,
			Msg:    "超过最大上限",
		}
	}
}

// Delete 删除购物车
func (service *CartService) Delete() serializer.Response {
	var cart model.Cart
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Update 修改购物车信息
func (service *CartService) Update() serializer.Response {
	var cart model.Cart
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cart.Num = service.Num
	err = model.DB.Save(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Show 订单
func (service *CartService) Show(id string) serializer.Response {
	var carts []model.Cart
	code := e.SUCCESS

	err := model.DB.Where("user_id=?", id).Find(&carts).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(carts),
	}
}
