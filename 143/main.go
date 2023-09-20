package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/check/health", func(ctx *gin.Context) {
		log.Println("ok")
	})

	r.Run(":8080")
}
