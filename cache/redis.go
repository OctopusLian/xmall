/*
 * @Author: neozhang
 * @Date: 2022-06-09 16:22:34
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-09 22:34:06
 * @Description: 请填写简介
 */
package cache

import (
	"strconv"
	"xmall/logging"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint("0", 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		logging.Info(err)
		panic(err)
	}

	RedisClient = client
}
