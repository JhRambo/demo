package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/random_string", func(c *gin.Context) {
		randomString := generateRandomString()
		c.String(http.StatusOK, randomString)
	})

	r.Run(":8080")
}

func generateRandomString() string {
	// 使用当前时间作为随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机字符串
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
