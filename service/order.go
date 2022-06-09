/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:20:11
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 16:28:19
 * @Description: 请填写简介
 */
package service

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"xmall/cache"
	"xmall/e"
	"xmall/logging"
	"xmall/model"
	"xmall/serializer"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type OrderService struct {
	//添加订单
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressID uint `form:"address_id" json:"address_id"`
	//查询订单列表
	Limit int  `form:"limit" json:"limit"`
	Start int  `form:"start" json:"start"`
	Type  uint `form:"type" json:"type" `
}

// Create 创建订单
func (service *OrderService) Create() serializer.Response {
	order := model.Order{
		UserID:    service.UserID,
		ProductID: service.ProductID,
		Num:       service.Num,
		Type:      1,
	}
	address := model.Address{}
	code := e.SUCCESS
	//查找对应的地址
	if err := model.DB.First(&address, service.AddressID).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressName = address.Name
	order.AddressPhone = address.Phone
	order.Address = address.Address
	//生成随机订单号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(service.UserID))
	number = number + productNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		logging.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.OrderNum = orderNum
	//存入数据库
	err = model.DB.Create(&order).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//将订单号存入Redis,并设置过期时间
	data := redis.Z{Score: float64(time.Now().Unix()) + 15*time.Minute.Seconds(), Member: orderNum}
	cache.RedisClient.ZAdd(os.Getenv("REDIS_ZSET_KEY"), data)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) List(id string) serializer.Response {
	var orders []model.Order

	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 5
	}

	if service.Type == 0 {
		if err := model.DB.Model(&orders).Where("user_id=?", id).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(&orders).Where("user_id=? AND type=?", id, service.Type).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("user_id=? AND type=?", id, service.Type).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}

func (service *OrderService) Show(num string) serializer.Response {
	var order model.Order
	var product model.Product
	code := e.SUCCESS
	//根据id查找order
	if err := model.DB.Where("order_num=?", num).First(&order).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//根据order查找product
	if err := model.DB.Where("id=?", order.ProductID).First(&product).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_PRODUCT
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product),
	}
}
