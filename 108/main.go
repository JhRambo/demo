package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379", // Redis 服务器地址和端口
		Password: "",                // 密码，如果有的话
		DB:       0,                 // 数据库编号
	})

	// 定义 Lua 脚本
	script := `
		local result = redis.call("GET", KEYS[1])
		if result ~= nil then
			redis.call("EXPIRE", KEYS[1], ARGV[1])
		end
		return result
	`

	// 定义键和过期时间
	key := "lua"
	expiration := 3600

	// 执行 Lua 脚本
	result, err := client.Eval(context.TODO(), script, []string{key}, expiration).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("Error:", err)
		return
	}

	if err == redis.Nil {
		fmt.Println("Key does not exist")
		return
	}

	fmt.Println("Result:", result)
}
