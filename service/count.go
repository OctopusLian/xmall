/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 18:47:24
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 19:02:04
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type ShowCountService struct {
}

func (service *ShowCountService) Show(id string) serializer.Response {
	code := e.SUCCESS
	var favoriteTotal int
	var notPayTotal int
	var payTotal int
	if err := model.DB.Model(model.Favorite{}).Where("user_id=?", id).Count(&favoriteTotal).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Model(model.Order{}).Where("user_id=? AND type=?", id, 1).Count(&notPayTotal).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Model(model.Order{}).Where("user_id=? AND type=?", id, 2).Count(&payTotal).Error; err != nil {
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
		Data:   serializer.BuildCount(favoriteTotal, notPayTotal, payTotal),
	}
}
