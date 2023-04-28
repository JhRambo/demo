package main

import (
	"encoding/json"
	"fmt"
)

var MapZdy = map[string]string{} //全局变量 设置一个map类型的变量，用于设置自定义context 参数key的值

func main() {
	MapZdyData("testinfo", []string{"1", "2", "3"})
	MapZdyData("userinfo", []string{"a", "b", "c"})
	fmt.Println("MapZdy==========", MapZdy)
}

func MapZdyData(str string, x interface{}) map[string]string {
	u, _ := json.Marshal(x)
	MapZdy[str] = string(u)
	return MapZdy
}
