/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:26:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:26:23
 */
package serializer

import "xmall/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		//Avatar:    user.AvatarURL(),
		CreatedAt: user.CreatedAt.Unix(),
	}
}
