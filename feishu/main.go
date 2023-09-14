package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 计算飞书机器人请求签名
func getFeishuRobotRequestSign(timestamp int64, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

type FeishuMessage struct {
	Timestamp string           `json:"timestamp"`
	Sign      string           `json:"sign"`
	MsgType   string           `json:"msg_type"`
	Content   FeishuMsgContent `json:"content"`
}

type FeishuMsgContent struct {
	Text string `json:"text"`
}

// 处理推送飞书机器人的请求
func pushFeishu(c *gin.Context) {
	// 读取请求体数据
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	var data FeishuMessage
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	// 飞书机器人配置
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/99593ee7-3d2d-4f06-8910-3e6239ee33a3"
	secret := "HSYRr2RBdXqznHV08AIEQe"

	// 计算当前时间戳
	timestamp := time.Now().Unix()
	// 计算签名
	sign, err := getFeishuRobotRequestSign(timestamp, secret)

	data.Timestamp = fmt.Sprint(timestamp)
	data.Sign = sign

	// 将数据转换为 JSON 字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sign"})
	}

	// 发送 HTTP POST 请求
	client := http.Client{}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// 返回响应
	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": string(bodyResponse)})
}

func main() {
	// 创建 Gin 引擎
	router := gin.Default()

	// 配置路由
	router.POST("/pushFeishu", pushFeishu)

	// 启动服务器
	router.Run(":8080")
}
