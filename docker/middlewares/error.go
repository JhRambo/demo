package middlewares

import (
	"demo/docker/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(ctx *gin.Context) {
	ctx.Next()

	status := ctx.Writer.Status()
	if status >= http.StatusInternalServerError {
		errorMessage := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		logError(status, ctx.Request.Method, ctx.Request.URL.Path, errorMessage)
	}
}

func logError(status int, method string, path string, errorMessage string) {
	now := time.Now()
	logDir := "logs"
	logFileName := fmt.Sprintf("error-%s.log", now.Format("20060102"))
	logFilePath := filepath.Join(logDir, logFileName)

	utils.CreateDir(logDir)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	log.SetOutput(file)

	// 记录错误的相关信息
	log.Printf("Status: %d", status)
	log.Printf("Method: %s", method)
	log.Printf("Path: %s", path)
	log.Printf("Message: %s", errorMessage)
}
