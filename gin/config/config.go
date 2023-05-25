package config

import "fmt"

const SERVER_NAME1 = "grpc-server-1"
const SERVER_PORT = 8081 //grpc-server端口
const GW_PORT = 8088     //gateway网关端口

var Grpcserver = map[string]string{
	SERVER_NAME1: fmt.Sprintf(":%d", SERVER_PORT),
}

type GWResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
