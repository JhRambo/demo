package config

import "time"

const (
	TOKENS_PER_SECOND  = 999              //每秒令牌数量	TODO
	TOKENS_BUCKET_SIZE = 999              //桶令牌总量	TODO
	IP_RATE_MINUTE     = 999              //同一个ip每分钟允许的最大请求次数	TODO
	TIMEOUT            = time.Second * 30 //请求超时时间	TODO
)

type GWResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
