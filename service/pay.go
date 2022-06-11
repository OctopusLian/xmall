/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 19:18:32
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 19:22:04
 */
package service

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"xmall/cache"
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"
)

type PayService struct {
	OrderNum string `form:"order_num" json:"order_num" `
	PayType  string `form:"pay_type" json:"pay_type" `
	Amount   string `form:"amount" json:"amount"`
	//接收FM支付回调接口
	MerchantNum     string `form:"merchantNum" json:"merchantNum" `
	OrderNo         string `form:"orderNo" json:"orderNo" `
	PlatformOrderNo string `form:"platformOrderNo" json:"platformOrderNo"`
	ActualPayAmount string `form:"actualPayAmount" json:"actualPayAmount" `
	State           int    `form:"state" json:"state" `
	Attch           string `form:"attch" json:"attch" `
	PayTime         string `form:"payTime" json:"payTime" `
	Sign            string `form:"sign" json:"sign" `
}

type PayOrderInfo struct {
	ID     string `json:"id"`     //渠道唯一ID
	PayURL string `json:"payUrl"` //支付页URL
}

type Result struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data PayOrderInfo `json:"data"`
}

func (service *PayService) Init() serializer.Response {
	code := e.SUCCESS
	//计算签名
	var buff bytes.Buffer
	buff.WriteString(os.Getenv("FM_Pay_ID"))
	buff.WriteString(service.OrderNum)
	buff.WriteString(service.Amount)
	buff.WriteString(os.Getenv("FM_Pay_NotifyURL"))
	buff.WriteString(os.Getenv("FM_Pay_Key"))
	sign := fmt.Sprintf("%x", md5.Sum(buff.Bytes()))

	returnURL := os.Getenv("FM_Pay_ReturnURL") + service.OrderNum
	//构造请求参数
	buff.Reset()
	buff.WriteString("sign=")
	buff.WriteString(sign)
	buff.WriteString("&amount=")
	buff.WriteString(service.Amount)
	buff.WriteString("&orderNo=")
	buff.WriteString(service.OrderNum)
	buff.WriteString("&payType=")
	buff.WriteString(service.PayType)
	buff.WriteString("&merchantNum=")
	buff.WriteString(os.Getenv("FM_Pay_ID"))
	buff.WriteString("&notifyUrl=")
	buff.WriteString(os.Getenv("FM_Pay_NotifyURL"))
	buff.WriteString("&returnUrl=")
	buff.WriteString(returnURL)
	buff.WriteString("&attch=")
	buff.WriteString(os.Getenv("FM_Pay_attch"))
	//调用渠道接口
	resp, err := http.Post("http://zfapi.nnt.ltd/api/startOrder", "application/x-www-form-urlencoded", &buff)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_CALL_API
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//读取所有响应数据
	var data []byte
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		logging.Info(err)
		code = e.ERROR_READ_FILE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//解析渠道返回的JSON
	var result Result
	if err = json.Unmarshal(data, &result); err != nil {
		logging.Info(err)
		code = e.ERROR_UNMARSHAL_JSON
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: result.Code,
		Msg:    result.Msg,
		Data:   result.Data,
	}
}

func (service *PayService) Confirm() {
	if service.Attch == os.Getenv("FM_Pay_attch") {
		if service.State == 1 {
			if err := model.DB.Model(model.Order{}).Where("order_num=?", service.OrderNo).Update("type", 2).Error; err != nil {
				logging.Info(err)
			}
			if err := cache.RedisClient.ZRem(os.Getenv("REDIS_ZSET_KEY"), service.OrderNo).Err(); err != nil {
				logging.Info(err)
			}
		}
	}
}
