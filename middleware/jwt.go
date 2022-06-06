/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-07 07:08:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-07 07:09:55
 */
package middleware

import (
	"time"
	"xmall/status"
	"xmall/util"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = status.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = status.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != status.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    status.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
