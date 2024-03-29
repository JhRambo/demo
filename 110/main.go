package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserInfo struct {
	Name string
	Age  int
}

func (p *UserInfo) MarshalBinary() ([]byte, error) {
	// 编组结构体为二进制数据
	// 在这个例子中，我们使用 JSON 编码
	return json.Marshal(p)
}

func (p *UserInfo) UnmarshalBinary(data []byte) error {
	// 解码二进制数据为结构体
	// 在这个例子中，我们使用 JSON 解码
	return json.Unmarshal(data, p)
}

var rdb *redis.Client
var ttl = time.Duration(1800) * time.Second

func init() {
	// 创建 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "", // 如果没有密码，则设置为空字符串
		DB:       0,  // Redis 默认数据库
	})

}

func main() {
	// 创建结构体实例
	person := &UserInfo{
		Name: "Alice",
		Age:  33,
	}

	// // 将结构体转换为 JSON 字符串
	// jsonData, err := json.Marshal(person)
	// if err != nil {
	// 	fmt.Println("JSON encoding error:", err)
	// 	return
	// }

	ctx := context.Background()
	key := "persons"

	// 存储结构体到 Redis Set
	err := rdb.Set(ctx, key, person, ttl).Err()
	if err != nil {
		fmt.Println("Redis Set error:", err)
		return
	}

	fmt.Println("Struct stored in Redis Set successfully.")

	val, _ := GetUserInfo(ctx, key)
	fmt.Println(val)
}

func GetUserInfo(ctx context.Context, key string) (*UserInfo, error) {
	var userInfo UserInfo
	err := rdb.Get(ctx, key).Scan(&userInfo)
	if err != nil {
		return nil, err
	}
	rdb.Expire(ctx, key, ttl)
	return &userInfo, nil
}
