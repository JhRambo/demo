package main

import (
	"fmt"
	"net/http"
	"strconv"

	"demo/jwt/middleware"
	"demo/jwt/model"
	"demo/jwt/token"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 通过鉴权账户,并生成对应的token进行返回
func authHandler(c *gin.Context) {
	user := &model.UserInfo{}
	err := c.ShouldBindJSON(user) //json数据格式绑定
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "msg": "invalid params"})
		return
	}
	// 检查账号是否存在,并生成一个token
	if user.UserName == "go" && user.PassWord == "123456" {
		tokenString, _ := token.GetToken(user.UserName)
		c.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{"token": tokenString}})
		return
	}
	c.JSON(200, gin.H{"code": 2002, "msg": "鉴权失败"})
	return
}

func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

// ssh加密验证
func TSSHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("start ->")
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			fmt.Println("err = ", err.Error())
			return
		}
		fmt.Println("+ err = ", err)
		c.Next()
		fmt.Println("success")
	}
}

func main() {
	r := gin.Default()
	// r.Use(TSSHandler(8088))
	r.POST("/auth", authHandler)
	r.GET("/home", middleware.Auth{}.JwtAuthMiddleware(), homeHandler) //中间件鉴权
	fmt.Println("ser is running")
	r.Run(":8088") // http
	// r.RunTLS(":8080", "key/cert.pem", "key/key.pem") // https
}
