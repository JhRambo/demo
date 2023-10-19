package utils

import (
	"demo/docker/config"
	"fmt"
)

// 格式化错误信息
func GetError(str string) error {
	return fmt.Errorf(str)
}

// 获取响应结构体
func GetResponse(codeValue int32) *config.GWResponse {
	return &config.GWResponse{}
}
