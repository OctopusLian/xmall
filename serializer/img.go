/*
 * @Author: neozhang
 * @Date: 2022-06-09 15:40:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 15:40:09
 * @Description: 请填写简介
 */
package serializer

import "xmall/model"

// ProductImg 商品图片序列化器
type ProductImg struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildImg 序列化商品图片
func BuildImg(item model.ProductImg) ProductImg {
	return ProductImg{
		ID:        item.ID,
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildImgs 序列化商品图片列表
func BuildImgs(items []model.ProductImg) (imgs []ProductImg) {
	for _, item := range items {
		img := BuildImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}
