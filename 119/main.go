package main

import (
	"fmt"
	"reflect"
)

type Fields struct {
	Platform   string `json:"platform" label:"平台"`
	ErrorInfo  string `json:"errorInfo" label:"报错内容"`
	UserInfo   string `json:"userInfo" label:"登录信息"`
	DeviceInfo string `json:"deviceInfo" label:"设备信息"`
	Time       string `json:"time" label:"时间"`
	ServerInfo string `json:"serverInfo" label:"服务器信息"`
}

func main() {

	var reqData = make(map[string]interface{}, 0)

	fields := Fields{
		Platform:   "1",
		ErrorInfo:  "2",
		UserInfo:   "3",
		DeviceInfo: "4",
		Time:       "5",
		ServerInfo: "6",
	}

	fmt.Println(fields)

	t := reflect.TypeOf(fields)
	v := reflect.ValueOf(fields)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		label := field.Tag.Get("label")
		value := v.Field(i).Interface()
		reqData[label] = value
	}
	fmt.Println(reqData)
}
