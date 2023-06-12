package main

import (
	"log"

	"demo/96/config"
	"demo/96/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := utils.Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		log.Fatal(err)
	}

	// //1.新增
	// jsonData := config.JsonData
	// var data interface{}
	// err = json.Unmarshal([]byte(jsonData), &data)
	// _, err = utils.InsertOne(config.COLLECTION, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //2.更新
	// filter := bson.M{"docId": "100"}
	// update := bson.M{"$set": bson.M{"data.node1.room.size": "200"}}
	// _, err = utils.UpdateOne(config.COLLECTION, filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //3.删除
	// filter := bson.M{"docId": "100"}
	// _, err = utils.DeleteOne(config.COLLECTION, filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //4.查询
	// filter := bson.M{"docId": "100"}
	// // filter := bson.M{"$and": []bson.M{
	// // 	{"data.node1.room.size": "200"},
	// // 	{"data.node1.room.style": "古风"},
	// // }}
	// var result bson.M
	// utils.FindOne(config.COLLECTION, filter, &result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Found JSON data:", result)

	// //5.高级用法
	// filter := bson.M{"$and": []bson.M{
	// 	{"docId": "100"},
	// 	{"data.node1.room.goods": bson.M{"$exists": true}},
	// }}
	// update := bson.M{"$set": bson.M{"data.node1.room.goods": "computer"}}
	// result, err := utils.UpdateOne(config.COLLECTION, filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 判断更新结果
	// if result.MatchedCount == 0 {
	// 	// 如果文档不存在，则插入新字段
	// 	_, err = utils.UpdateOne(config.COLLECTION, bson.M{"docId": "100"}, bson.M{"$set": bson.M{"data.node1.room.goods": "computer"}})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("Inserted new field!")
	// } else {
	// 	fmt.Println("Updated field!")
	// }

	owner := "100.1000" //锁拥有者
	// 6.分布式锁
	lockRes, err := utils.SetLock("node1", owner, int64(config.LOCKTTL))
	if err != nil {
		log.Fatal(err)
	}

	if lockRes["owner"] != owner { //判断是否是当前owner
		ok, err := utils.CheckLock("node1")
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			//删除成功
			utils.SetLock("node1", owner, int64(config.LOCKTTL))
			goto A
		}
		log.Fatal("node1 is already locked by another client")
	} else {
		goto A
	}

A:
	// time.Sleep(100 * time.Second)
	// 构造修改文档
	filter := bson.M{"docId": "100", "userId": "1000"}
	update := bson.M{
		"$set": bson.M{
			"data.node1.room.style": "乡村300",
		},
	}
	// 执行修改操作
	_, err = utils.UpdateOne(config.COLLECTION, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭连接
	defer utils.Disconnect()

	// 释放锁
	defer utils.UnLock("node1", owner)
}
