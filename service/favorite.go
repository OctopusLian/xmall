/*
 * @Author: neozhang
 * @Date: 2022-06-09 15:46:11
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:49:36
 * @Description: 请填写简介
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type FavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Limit     int  `form:"limit"`
	Start     int  `form:"start"`
}

func (service *FavoriteService) Create() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS
	model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite)
	if favorite == (model.Favorite{}) {
		favorite = model.Favorite{
			UserID:    service.UserID,
			ProductID: service.ProductID,
		}
		if err := model.DB.Create(&favorite).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ERROR_EXIST_FAVORITE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Delete() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&favorite).Error
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

func (service *FavoriteService) Show(id string) serializer.Response {
	var favorites []model.Favorite
	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 12
	}
	if err := model.DB.Model(&favorites).Where("user_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&favorites).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(total))
}
