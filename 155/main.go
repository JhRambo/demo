package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostRequest(url string, body interface{}, authorizationHeader string) (*http.Response, error) {
	// 将请求体转换为 JSON 格式
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 设置请求头部为 JSON 类型
	req.Header.Set("Content-Type", "application/json")

	// 添加可选的授权头部字段
	if authorizationHeader != "" {
		req.Header.Set("Authorization", authorizationHeader)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {
	// 示例使用方法
	url := "http://192.168.10.103:38101/v2/space/res/get" // 替换为实际的请求 URL
	body := map[string]interface{}{
		"configId": 1,
		"level":    3,
	} // 替换为实际的请求体数据
	authorizationHeader := "Bearer 222" // 替换为实际的授权头部字段（如果需要）

	resp, err := PostRequest(url, body, authorizationHeader)
	if err != nil {
		fmt.Println("请求发生错误:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	// ...
	// 读取响应结果
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("无法解析响应结果:", err)
		return
	}
	fmt.Println(result)
}
