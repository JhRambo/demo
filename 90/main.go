package main

import (
	"encoding/json" //文件位置：D:\MySoft\Go\src\encoding\json\stream.go
	"fmt"
	"strings"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	//错误的json格式
	jsonStr := `[
        {"id":1, "name":"Alice"},
        {"id":2, "name":"Bob"},
        {"id":3, "name":"Charlie"},
        {"id":4, "name":"David"},
        {"id":5, "name":"Eva"},
        {"id":6, "name":"Frank"},
    ]`

	dec := json.NewDecoder(strings.NewReader(jsonStr))
	for {
		var user User
		if err := dec.Decode(&user); err != nil {
			if err.Error() == "EOF" {
				break
			}
			if e, ok := err.(*json.SyntaxError); ok && e.Offset >= 0 {
				fmt.Printf("JSON语法错误：%s", jsonStr[e.Offset-10:e.Offset+10])
			} else {
				fmt.Println("解析错误：", err)
			}
			return
		}
		fmt.Println(user)
	}
}
