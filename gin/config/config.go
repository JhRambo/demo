package config

import (
	"fmt"
)

const SERVER_NAME1 = "grpc-server-1" //grpc服务名
const SERVER_PORT = 8081             //grpc-server端口
const GW_PORT = 8088                 //gateway网关端口

var GRPCserver = map[string]string{
	SERVER_NAME1: fmt.Sprintf(":%d", SERVER_PORT),
}

type GWResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
