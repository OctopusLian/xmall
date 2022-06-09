/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:56:27
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:57:32
 * @Description: 请填写简介
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type CarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

func (service *CarouselService) Create() serializer.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS

	err := model.DB.Create(&carousel).Error
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
		Data:   serializer.BuildCarousel(carousel),
	}
}

func (service *CarouselService) List() serializer.Response {
	carousels := []model.Carousel{}
	code := e.SUCCESS

	if err := model.DB.Find(&carousels).Error; err != nil {
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
		Data:   serializer.BuildCarousels(carousels),
	}
}
