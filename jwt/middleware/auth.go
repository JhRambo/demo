package middleware

import (
	"demo/jwt/token"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

// 中间件,认证token合法性
func (c Auth) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("authorization")
		if authHandler == "" {
			c.JSON(200, gin.H{"code": 2003, "msg": "请求头部auth为空"})
			c.Abort()
			return
		}
		jwt := strings.Split(authHandler, ".")
		cnt := 0
		for _, val := range jwt {
			cnt++
			if cnt == 3 {
				break
			}
			msg, _ := base64.StdEncoding.DecodeString(val)
			fmt.Println("val ->", string(msg))
		}
		// 我们使用之前定义好的解析JWT的函数来解析它,并且在内部解析时判断了token是否过期
		mc, err := token.ParseToken(authHandler)
		if err != nil {
			fmt.Println("err = ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.UserName)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
