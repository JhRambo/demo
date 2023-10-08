package main

import (
	"encoding/json"
	"fmt"
)

// 结构体嵌套
type Content struct {
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func main() {
	jsonData := `{"content":{"text":"{\"errorInfo\":{\"code\":-1,\"message\":\"出错了\",\"id\":123}}"}}`
	// jsonData := `{"content":{"text":"1111111111111111111122222222333bcaaaa你好"}}`
	var content Content
	err := json.Unmarshal([]byte(jsonData), &content)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return
	}

	fmt.Println(content)

	var response interface{}
	err = json.Unmarshal([]byte(content.Content.Text), &response)
	if err != nil {
		// 直接输出字符串
		fmt.Println(content.Content.Text)
		// fmt.Println("解析 response 出错:", err)
		return
	}

	fmt.Println(response)

	formattedJSON, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		fmt.Println("JSON 格式化出错:", err)
		return
	}

	fmt.Println(string(formattedJSON))
}
