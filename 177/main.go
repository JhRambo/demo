package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := "[{\"id\":\"510014223808466001\"},{\"id\":\"510014223808466001\"}]"
	var data interface{}
	json.Unmarshal([]byte(str), &data)
	jsonData := FormatToJson(data)
	fmt.Println(len(jsonData))
	fmt.Println(jsonData)
}

// 格式化json类型
func FormatToJson(body interface{}) string {
	// 将请求体转换为 JSON 格式
	jsonData, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
