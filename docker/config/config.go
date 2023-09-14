package config

import "time"

const (
	TOKENS_PER_SECOND  = 1               //每秒令牌数量	TODO
	TOKENS_BUCKET_SIZE = 1               //桶令牌总量	TODO
	TIMEOUT            = time.Second * 1 //超时时间
)

type GWResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
