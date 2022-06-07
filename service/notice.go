/*
 * @Author: neozhang
 * @Date: 2022-06-07 18:01:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 18:03:15
 * @Description: 请填写简介
 */
package service

import (
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type NoticeService struct {
	NoticeID uint   `form:"notice_id" json:"notice_id"`
	Text     string `form:"text" json:"text"`
}

// Create 公告创建的服务
func (service *NoticeService) Create() serializer.Response {
	notice := model.Notice{
		Text: service.Text,
	}
	code := e.SUCCESS
	err := model.DB.Create(&notice).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR_DATABASE
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

// Update 公告更新的服务
func (service *NoticeService) Update() serializer.Response {
	var notice model.Notice
	code := e.SUCCESS
	if err := model.DB.First(&notice, service.NoticeID).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	notice.Text = service.Text
	if err := model.DB.Save(&notice).Error; err != nil {
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

// Show 公告详情服务
func (service *NoticeService) Show() serializer.Response {
	var notice model.Notice
	code := e.SUCCESS
	if err := model.DB.First(&notice, 1).Error; err != nil {
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
		Data:   serializer.BuildNotice(notice),
	}
}
