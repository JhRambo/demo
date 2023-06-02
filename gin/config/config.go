package config

import (
	"fmt"
)

const SERVER_NAME1 = "grpc-server-1" //grpc服务名
const SERVER_PORT1 = 8081            //grpc-server端口
const GW_PORT = 8088                 //gateway网关端口

var GRPCserver = map[string]string{
	SERVER_NAME1: fmt.Sprintf(":%d", SERVER_PORT1),
}

type GWResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

const MSGPACK_URI = "/HELLO/SAYHELLO" //使用msgpack协议的接口uri
