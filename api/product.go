/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-08 22:15:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-08 22:23:10
 */
package api

import (
	"xmall/logging"
	"xmall/service"

	"github.com/gin-gonic/gin"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	service := service.ProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListProducts 商品列表接口
func ListProducts(c *gin.Context) {
	service := service.ProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowProduct 商品详情接口
func ShowProduct(c *gin.Context) {
	service := service.ProductService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// DeleteProduct 删除商品的接口
func DeleteProduct(c *gin.Context) {
	service := service.ProductService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// UpdateProduct 更新商品的接口
func UpdateProduct(c *gin.Context) {
	service := service.ProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	service := service.ProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Searchs()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
