package utils

import (
	"demo/96/config"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
)

// 更新owner
func UpdateOwner() bool {
	return true
}

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

// 初始化对象资源
func Create() (*config.Response, error) {
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

	jsonData := config.JsonData
	var data interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	_, err = InsertOne(config.COLLECTION, data)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}

	keys := bson.D{
		bson.E{Key: "configId", Value: 1},
	}
	// 创建索引
	err = CreateIndexes(config.COLLECTION, keys, true)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	resp.Message = "ok"
	return resp, nil
}

/*
	更新空间资源

owner锁拥有者 eid
*/
func Update(configId int64, data []map[string]interface{}, owner int64) (*config.Response, error) {
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

	logs := make([]map[string]interface{}, 0)
	for _, v := range data {
		upData := v["data"]
		var path string
		if value, ok := v["path"].(string); ok {
			path = value
		}
		paths, err := OperateStr(path)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
		operateType := paths[0] //操作类型，目前只有两种nodeList basedata
		if operateType == "nodeList" {
			nodeInfo := fmt.Sprintf("%v.$.%v", operateType, "baseInfo.name")
			filter := bson.M{"configId": configId, "eid": owner, "nodeList.id": v["id"]}
			update := bson.M{"$set": bson.M{nodeInfo: v["data"]}}
			UpdateOne(config.COLLECTION, filter, update)
			return nil, nil
		} else {

		}

		lockPath := ""
		if len(paths) > 1 {
			// 6.分布式锁
			lockPath = paths[0] + "." + paths[1]
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
					//删除成功，重新加锁
					SetLock(lockPath, owner, int64(config.LOCKTTL))
					goto A
				}
				return resp, fmt.Errorf("current path is already locked by another client")
			} else {
				goto A
			}
		}
	A:
		// 模拟请求耗时操作
		// time.Sleep(30 * time.Second)
		// 构造查询条件
		filter := bson.M{"$and": []bson.M{
			{"configId": configId, "eid": owner},
			{path: bson.M{"$exists": true}},
		}}
		update := bson.M{}
		if v["action"] == "d" { //删除
			if _, err := strconv.Atoi(paths[len(paths)-1]); err == nil { //判断path最后一个字符串是否是整型数字，如果是则说明当前操作的是数组
				// 删除子数组========================================
				path_p := ""
				lastDotIndex := strings.LastIndex(path, ".")
				if lastDotIndex > 0 {
					path_p = path[:lastDotIndex]
				} else {
					if len(paths) > 1 {
						// 释放锁
						UnLock(lockPath, owner)
					}
					resp.Message = err.Error()
					return resp, err
				}
				_, err = UpdateOne(config.COLLECTION, filter, bson.M{
					"$unset": bson.M{
						path: 1,
					}})
				if err != nil {
					if len(paths) > 1 {
						// 释放锁
						UnLock(lockPath, owner)
					}
					resp.Message = err.Error()
					return resp, err
				}
				update = bson.M{"$pull": bson.M{
					path_p: nil,
				}}
			} else {
				update = bson.M{
					"$unset": bson.M{
						path: 1,
					},
				}
			}
		} else if v["action"] == "e" { //数组调整顺序
			var slices bson.M
			// 构造聚合管道
			pipeline := []bson.M{
				{
					"$match": bson.M{"configId": configId, "eid": owner},
				},
				{
					"$project": bson.M{
						path:  1,
						"_id": 0,
					},
				},
			}
			err = Aggregate(config.COLLECTION, pipeline, slices)
			if err != nil {
				if len(paths) > 1 {
					// 释放锁
					UnLock(lockPath, owner)
				}
				resp.Message = err.Error()
				return resp, err
			}
		} else {
			update = bson.M{
				"$set": bson.M{
					path: upData,
				},
			}
		}
		// 存在即更新，不存在即插入================================
		actionType := "" //操作类型
		result, err := UpdateOne(config.COLLECTION, filter, update)
		if err != nil {
			if len(paths) > 1 {
				// 释放锁
				UnLock(lockPath, owner)
			}
			resp.Message = err.Error()
			return resp, err
		}
		// 判断更新结果
		if v["action"] == "d" { //删除
			if result.ModifiedCount == 0 {
				if len(paths) > 1 {
					// 释放锁
					UnLock(lockPath, owner)
				}
				resp.Message = "Not found field!"
				return resp, nil
			}
			actionType = "delete"
			resp.Message = "Deleted field!"
		} else {
			if result.MatchedCount > 0 {
				actionType = "update"
				resp.Message = "Updated field!"
			} else {
				// 如果不存在，则插入新字段
				_, err = UpdateOne(config.COLLECTION, bson.M{"configId": configId, "eid": owner}, update)
				if err != nil {
					if len(paths) > 1 {
						// 释放锁
						UnLock(lockPath, owner)
					}
					resp.Message = err.Error()
					return resp, err
				}
				actionType = "create"
				resp.Message = "Inserted field!"
			}
		}
		logs = append(logs, map[string]interface{}{
			"actionType": actionType,
			"typeId":     v["typeId"],
			"info":       fmt.Sprintf("%v:%v:%v", path, upData, actionType),
		})
		if len(paths) > 1 {
			// 释放锁
			UnLock(lockPath, owner)
		}
	}
	// 记录操作日志
	RecordLogs(configId, logs)
	key := fmt.Sprintf("configId_%d", configId)
	recordLogsNum := IncrLogs(key)
	if recordLogsNum > 0 && (recordLogsNum%config.RECORDLOGSNUM) == 0 {
		//复制一份最新的文档到新的集合中
		var curLog map[string]interface{}
		filter := bson.M{"configId": configId}
		projection := bson.M{"_id": 0} //过滤_id字段
		FindOneProjection(config.COLLECTION, filter, &curLog, projection)
		curLog["createAt"] = time.Now().Format("2006-01-02 15:04:05")
		_, err = InsertOne(config.COLLECTION_LOGS, curLog)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
	}

	return resp, nil
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

func makeSliceUpdater(path string, indexes []int) []interface{} {
	var updater []interface{}
	updater = append(updater, "$map")
	updater = append(updater, bson.M{
		"input": bson.M{"$literal": bson.A{bson.M{}}},
		"as":    "this",
		"in": bson.M{
			"$arrayElemAt": []interface{}{
				bson.M{
					"$arrayElemAt": []interface{}{fmt.Sprintf("$%v", path), "$$this"},
				},
				bson.M{
					"$indexOfArray": []interface{}{indexes, "$$this"},
				},
			},
		},
	})
	updater = append(updater, bson.M{"$sort": bson.M{"this": 1}})
	updater = append(updater, "$map")
	updater = append(updater, bson.M{
		"input": bson.M{"$range": []interface{}{0, len(indexes)}},
		"as":    "i",
		"in": bson.M{
			"$arrayElemAt": []interface{}{
				fmt.Sprintf("$%v", path),
				bson.M{"$arrayElemAt": []interface{}{"$$arr", "$$i"}},
			},
		},
	})
	return []interface{}{updater}
}
