package main

import (
	"context"
	pb "demo/utils/proto/student"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

func main() {
	// 建立 MongoDB 连接
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// 选择数据库和集合
	collection := client.Database("test").Collection("student")

	// 创建一个 Student 实例
	student := &pb.Student{
		Name:  "Alice",
		Id:    12345,
		Email: "alice@example.com",
	}

	// // 将 Protocol Buffers 数据序列化为 BSON 格式
	// bsonData, err := bson.Marshal(student)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 插入数据到 MongoDB
	// _, err = collection.InsertOne(context.Background(), bsonData)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 将 Protocol Buffers 数据转换为字节流
	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}
	// 将数据存储到 MongoDB 中
	_, err = collection.InsertOne(context.Background(), bson.M{"data": data})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted data into MongoDB")

	// 构造查询条件
	id, err := primitive.ObjectIDFromHex("658bb50504d7526832a29c4a")
	if err != nil {
		log.Fatal(err)
	}
	filter := primitive.D{{Key: "_id", Value: id}}

	// 查询数据
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		log.Fatal(result.Err())
	}

	// 从查询结果中解析出包含"data"字段的文档
	var document bson.M
	err = result.Decode(&document)
	if err != nil {
		log.Fatal(err)
	}
	// 获取嵌套文档中的"data"字段的值
	binaryData := document["data"].(primitive.Binary)
	// 从二进制数据中取出字节切片
	da := binaryData.Data

	// 反序列化二进制数据为 Protocol Buffers 数据
	stu := &pb.Student{}
	err = proto.Unmarshal(da, stu)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(stu)
}
