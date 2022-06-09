/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-08 22:26:40
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:19:34
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

//地址管理
type AddressService struct {
	ID        uint   `form:"id" json:"id"`
	UserID    uint   `form:"user_id" json:"user_id"`
	Name      string `form:"name" json:"name"`
	Phone     string `form:"phone" json:"phone"`
	Address   string `form:"address" json:"address"`
	AddressID uint   `form:"address_id" json:"address_id"`
}

func (service *AddressService) Create() serializer.Response {
	var address model.Address
	code := e.SUCCESS
	address = model.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}

	err := model.DB.Create(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []model.Address
	err = model.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}

func (service *AddressService) Delete() serializer.Response {
	var address model.Address
	code := e.SUCCESS

	err := model.DB.Where("id=?", service.AddressID).Find(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&address).Error
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

func (service *AddressService) Show(id string) serializer.Response {
	var addresses []model.Address
	code := e.SUCCESS

	err := model.DB.Where("user_id=?", id).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}

func (service *AddressService) Update() serializer.Response {
	address := model.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	address.ID = service.ID
	code := e.SUCCESS
	err := model.DB.Save(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []model.Address
	err = model.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}
