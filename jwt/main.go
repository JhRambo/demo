package main

import (
	"log"
	"net/http"

	"demo/jwt/middleware"
	"demo/jwt/model"
	"demo/jwt/token"

	"github.com/gin-gonic/gin"
)

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

func main() {
	r := gin.Default()
	r.POST("/auth", authHandler)
	md := &middleware.Auth{}
	r.GET("/home", md.JwtAuthMiddleware(), homeHandler) //中间件鉴权
	log.Println("ser is running")
	r.Run(":8088")
}
