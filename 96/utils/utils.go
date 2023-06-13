package utils

import (
	"demo/96/config"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 更新owner
func UpdateOwner() bool {
	return true
}

// 提取path路径
func OperateStr(path string) ([]string, error) {
	sliceStr := strings.Split(path, ".")
	if len(sliceStr) < 3 {
		return nil, fmt.Errorf("path格式不对，至少3级")
	}
	return sliceStr, nil
}

// 初始化对象资源
func Create() *config.Response {
	resp := &config.Response{
		Code: 0,
	}
	err := Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		log.Fatal(err)
	}
	// 关闭连接
	defer Disconnect()

	jsonData := config.JsonData
	var data interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	_, err = InsertOne(config.COLLECTION, data)
	if err != nil {
		log.Fatal(err)
	}

	keys := bson.D{
		bson.E{Key: "configId", Value: 1},
	}
	// 创建索引
	err = CreateIndexes(config.COLLECTION, keys, true)
	if err != nil {
		log.Fatal(err)
	}
	resp.Message = "ok"
	return resp
}

/*
	更新空间资源

owner锁拥有者 eid
path := "data.node1.baseInfo.type"
val := "造型"
*/
func UpdateSpace(path string, val interface{}, action string, configId int64, spaceId int64, owner int64) (*config.Response, error) {
	resp := &config.Response{
		Code: 0,
	}
	err := Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	// 关闭连接
	defer Disconnect()

	paths, err := OperateStr(path)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}

	// 6.分布式锁
	lockPath := paths[0] + "." + paths[1] + "." + paths[2]
	lockPath = fmt.Sprintf("%d"+"."+lockPath, configId)
	lockRes, err := SetLock(lockPath, owner, int64(config.LOCKTTL))
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	if lockRes["owner"] != owner { //判断是否是当前owner
		ok, err := CheckLock(lockPath)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
		if ok {
			//删除成功
			SetLock(lockPath, owner, int64(config.LOCKTTL))
			goto A
		}
		log.Fatal("node1 is already locked by another client")
	} else {
		goto A
	}
A:
	// 模拟请求耗时操作
	// time.Sleep(30 * time.Second)
	// 构造查询条件
	filter := bson.M{"$and": []bson.M{
		{"configId": configId, "spaceId": spaceId, "eid": owner},
		{path: bson.M{"$exists": true}},
	}}

	update := bson.M{}
	if action == "d" { //删除
		// 定义更新操作
		update = bson.M{
			"$unset": bson.M{
				path: "",
			},
		}
	} else {
		// 定义更新操作
		update = bson.M{
			"$set": bson.M{
				path: val,
			},
		}
	}
	// 存在即更新，不存在即插入================================
	edittype := 0
	result, err := UpdateOne(config.COLLECTION, filter, update)
	if err != nil {
		UnLock(lockPath, owner)
		Disconnect()
		resp.Message = err.Error()
		return resp, err
	}
	// 判断更新结果
	if action == "d" {
		if result.ModifiedCount == 0 {
			UnLock(lockPath, owner)
			Disconnect()
			resp.Message = "Not found field!"
			return resp, nil
		}
		edittype = 3
		resp.Message = "Deleted field!"
		return resp, nil
	} else {
		if result.MatchedCount > 0 {
			edittype = 2
			resp.Message = "Updated field!"
		} else {
			// 如果不存在，则插入新字段
			_, err = UpdateOne(config.COLLECTION, bson.M{"configId": configId, "spaceId": spaceId, "eid": owner}, update)
			if err != nil {
				UnLock(lockPath, owner)
				Disconnect()
				resp.Message = err.Error()
				return resp, err
			}
			edittype = 1
			resp.Message = "Inserted field!"
		}
	}
	// 记录操作日志
	content := map[string]interface{}{
		"nodename": "智慧屏",
		"nodetype": 1,
		"edittype": edittype,
	}
	RecordLogs(configId, content)
	key := fmt.Sprintf("configId_%d", configId)
	recordLogsNum := IncrLogs(key)
	if recordLogsNum > 0 && (recordLogsNum%config.RECORDLOGSNUM) == 0 {
		//赋值一份最新的文档到新的集合中
		var curLog map[string]interface{}
		filter := bson.M{"configId": configId}
		projection := bson.M{"_id": 0}
		FindOneProjection(config.COLLECTION, filter, &curLog, projection)
		curLog["createAt"] = time.Now().Format("2006-01-02 15:04:05")
		_, err = InsertOne(config.COLLECTION_LOGS, curLog)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
	}
	// 释放锁
	defer UnLock(lockPath, owner)
	return resp, nil
}

// 记录节点操作日志
func RecordLogs(configId int64, content map[string]interface{}) {
	recordlogsData := map[string]interface{}{
		"configId": configId,
		"content": map[string]interface{}{
			"nodename": content["nodename"],
			"nodetype": content["nodetype"],
			"edittype": content["edittype"],
			"createAt": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	_, err = InsertOne(config.RECORDLOGS, recordlogsData)
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
	result, err := rdb.Incr(key).Result()
	if err != nil {
		panic(err)
	}
	return result
}

type Object interface {
	// 任意的方法定义
}

func GetVal(rsMap map[string]interface{}, path string) interface{} {
	// 定义查询条件
	filterSelect := bson.M{}
	// 定义结果投影
	projection := bson.M{
		path:  1,
		"_id": 0,
	}
	var res interface{}
	FindOneProjection(config.COLLECTION, filterSelect, &res, projection)
	res, ok := res.(bson.D)
	if !ok {
		log.Fatal("Failed to cast query result to bson.D")
	}
	// 访问 data.node1.baseInfo.type 字段
	data, ok := rsMap["data"].(primitive.D)
	if !ok {
		log.Fatal("Failed to cast data field to primitive.D")
	}
	node1, ok := data.Map()["node1"].(primitive.D)
	if !ok {
		log.Fatal("Failed to cast node1 field to primitive.D")
	}
	baseInfo, ok := node1.Map()["baseInfo"].(primitive.D)
	if !ok {
		log.Fatal("Failed to cast baseInfo field to primitive.D")
	}
	fieldType := baseInfo.Map()["type"]
	switch fieldType.(type) {
	case string:
		//直接替换
		fmt.Println("The variable is an string")
	case Object:
		fmt.Println("The variable is an object")
	case []interface{}:
		fmt.Println("The variable is an []")
	}
	return fieldType
}
