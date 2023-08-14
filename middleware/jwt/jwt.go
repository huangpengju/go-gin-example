package jwt

import (
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 声明变量
		var code int
		var data interface{}

		// 赋值 200，表示ok
		code = e.SUCCESS
		// 获取 URL 参数 token
		token := c.Query("token")

		if token == "" {
			// 赋值 400 请求参数错误
			code = e.INVALID_PARAMS
		} else {
			// ParseToken 解析 Token
			claims, err := util.ParseToken(token)
			if err != nil {
				// 20001 表示Token鉴权失败
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 20002 表示Token已超时
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			// 401
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
