package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client
var err error

// 创建连接
func Connect(url string, dbName string) error {
	// 设置连接超时时间
	ctx, cancel := WithTimeout(10 * time.Second)
	defer cancel()

	// 建立连接
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	// 获取数据库
	DB = client.Database(dbName)

	return nil
}

// 关闭连接
func Disconnect() error {
	if client == nil {
		return nil
	}
	// 设置断开连接超时时间
	ctx, cancel := WithTimeout(10 * time.Second)
	defer cancel()

	return client.Disconnect(ctx)
}

// 获取集合
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}

// 设置超时时间
func WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// 新增
func InsertOne(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 更新
func UpdateOne(collectionName string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 删除
func DeleteOne(collectionName string, filter bson.M) (*mongo.DeleteResult, error) {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 查找数据
func FindOne(collectionName string, filter bson.M, result interface{}) error {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// 查找所有数据
func FindAll(collectionName string, filter bson.M, result interface{}) error {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

func FindOneAndUpdate(collectionName string, filter bson.M, update bson.M, result interface{}, options *options.FindOneAndUpdateOptions) error {
	collection := GetCollection(collectionName)
	ctx, cancel := WithTimeout(5 * time.Second)
	defer cancel()

	err := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
