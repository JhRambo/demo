package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := "[{\"id\":\"510014223808466001\"},{\"id\":\"510014223808466001\"}]"
	arr, _ := CheckCreatorData(str)
	for _, item := range arr {
		fmt.Println(item["id"])
	}
}

// 检查数据
func CheckCreatorData(data string) ([]map[string]interface{}, error) {
	var creatorData interface{}
	err := json.Unmarshal([]byte(data), &creatorData)
	if err != nil {
		return nil, err
	}
	jsonString, err := FormatToJson(creatorData)
	if err != nil {
		return nil, err
	}
	var datas []map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &datas)
	if err != nil {
		return nil, err
	}
	if len(datas) <= 0 {
		return nil, fmt.Errorf("参数错误")
	}
	return datas, nil
}

// 格式化json类型
func FormatToJson(body interface{}) (string, error) {
	// 将请求体转换为 JSON 格式
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
