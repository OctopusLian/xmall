/*
 * @Author: neozhang
 * @Date: 2022-06-07 18:04:32
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 18:04:33
 * @Description: 请填写简介
 */
package serializer

import "xmall/model"

// Notice 公告序列化器
type Notice struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNotice 序列化轮播图
func BuildNotice(item model.Notice) Notice {
	return Notice{
		ID:        item.ID,
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
