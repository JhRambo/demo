package mongo

import (
	"demo/grpc/config"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
)

// 提取path路径
func OperateStr(path interface{}) ([]string, error) {
	if v, ok := path.(string); ok {
		sliceStr := strings.Split(v, ".")
		if len(sliceStr) < 1 {
			return nil, fmt.Errorf("path不能为空")
		}
		return sliceStr, nil
	}
	return nil, nil
}

// 记录节点操作日志
func RecordLogs(configId int64, content interface{}) {
	recordlogsData := map[string]interface{}{
		"configId": configId,
		"createAt": time.Now().Format("2006-01-02 15:04:05"),
		"content":  content,
	}
	_, err := InsertOne(config.RECORDLOGS, recordlogsData)
	if err != nil {
		log.Fatal(err)
	}
	keys := bson.D{
		bson.E{Key: "configId", Value: 1},
	}
	// 创建索引
	err = CreateIndexes(config.RECORDLOGS, keys, false)
	if err != nil {
		log.Fatal(err)
	}
}

// num++
func IncrLogs(key string) int64 {
	// 创建 Redis 客户端，连接到指定的 Redis 实例
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.103:36379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 使用 INCR 命令将指定键的值自增 1
	result, _ := rdb.Incr(key).Result()
	return result
}
