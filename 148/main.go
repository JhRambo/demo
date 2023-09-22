package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// type reqData struct {
// 	Msg_type string `json:"msg_type"`
// 	Content  struct {
// 		Text interface{} `json:"text"`
// 	} `json:"content"`
// }

type reqData struct {
	Msg_type string       `json:"msg_type"`
	Content  *ContentData `json:"content"`
}

type ContentData struct {
	Text interface{} `json:"text"`
}

func HttpPost(url string, bys []byte) (*http.Response, error) {
	return http.Post(url, "application/json", bytes.NewBuffer(bys))
}

func main() {
	url := "http://192.168.10.103:38401/v2/alarm/feishu/notify1"

	// req := &reqData{
	// 	Msg_type: "text",
	// 	Content: struct {
	// 		Text interface{} `json:"text"`
	// 	}{Text: "888"},
	// }

	req := &reqData{
		Msg_type: "text",
		Content: &ContentData{
			Text: "999",
		},
	}

	// 创建 JSON 消息体
	jsonBody, _ := json.Marshal(req)

	// 发送 POST 请求
	resp, err := HttpPost(url, jsonBody)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 处理响应
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
}
