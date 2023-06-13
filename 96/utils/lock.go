package utils

import (
	"demo/96/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 获取分布式锁
func SetLock(lockName string, owner int64, lockTTL int64) (primitive.M, error) {
	fmt.Println("setlock ok!")
	// 构造查询和更新文档
	filter := bson.M{"_id": lockName, "owner": owner}
	update := bson.M{"$set": bson.M{"locked": true, "owner": owner, "expireAt": time.Now().Unix() + lockTTL}}

	// 使用 findAndModify 命令获取锁，原子操作，一次操作完成查询和更新
	var result bson.M
	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	if err := FindOneAndUpdate(config.COLLECTION, filter, update, &result, options); err != nil {
		if err == mongo.ErrNoDocuments {
			// 没有符合条件的文档，说明锁已经被其他客户端获取
			return nil, fmt.Errorf("failed to get lock: lock has been taken by other client")
		}
		// 更新失败
		return nil, fmt.Errorf("failed to get lock: %v", err)
	}
	// 执行成功
	return result, nil
}

// 释放分布式锁
func UnLock(lockName string, owner int64) error {
	fmt.Println("unlock ok!")
	// 构造过滤器和更新器
	filter := bson.M{"_id": lockName, "locked": true, "owner": owner}

	// 执行删除锁操作
	result, err := DeleteOne(config.COLLECTION, filter)
	if err != nil {
		return err
	}
	// 判断删除是否成功
	if result.DeletedCount == 0 {
		return fmt.Errorf("lock %s not exists or already released", lockName)
	}

	return nil
}

// 检测锁是否已过期，过期删除
func CheckLock(lockName string) (bool, error) {
	filter := bson.M{"_id": lockName}
	var result bson.M
	err := FindOne(config.COLLECTION, filter, &result)
	if err != nil {
		return false, err
	}
	expireAt, ok := result["expireAt"].(int64)
	if !ok {
		return false, fmt.Errorf("failed to convert expireAt to string: %v", result["expireAt"])
	}
	curTime := time.Now().Unix()
	if curTime > expireAt {
		// 锁已经过期，直接删除锁
		_, err = DeleteOne(config.COLLECTION, filter)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
