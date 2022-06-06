/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:27:23
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:27:28
 */
package serializer

import "xmall/model"

// Admin 用户序列化器
type Admin struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAdmin 序列化用户
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:       admin.ID,
		UserName: admin.UserName,
		//Avatar:    admin.AvatarURL(),
		CreatedAt: admin.CreatedAt.Unix(),
	}
}
