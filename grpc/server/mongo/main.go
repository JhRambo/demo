package main

import (
	"context"
	"demo/grpc/config"
	pb "demo/grpc/proto/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"demo/grpc/utils/mongo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const server_port = 8081 //grpc-server端口
const gw_port = 8088     //gateway网关端口

type Server struct {
	pb.UnimplementedMongoHttpServer
}

func NewServer() *Server {
	return &Server{}
}

// 初始化对象资源 ok
func (s *Server) Create(ctx context.Context, req *pb.CreateHttpRequest) (*pb.CreateHttpResponse, error) {
	resp := &pb.CreateHttpResponse{
		Code: 0,
	}
	err := mongo.Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	// 关闭连接
	defer mongo.Disconnect()

	// jsonString := config.JsonData
	jsonString := req.Data //实际请求

	// fmt.Println(jsonString)
	// return nil, nil

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	data["configId"] = req.ConfigId
	data["spaceId"] = req.SpaceId
	data["eid"] = req.Eid
	_, err = mongo.InsertOne(config.COLLECTION, data)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}

	keys := bson.D{
		bson.E{Key: "configId", Value: 1},
	}
	// 创建索引
	err = mongo.CreateIndexes(config.COLLECTION, keys, true)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}

	resp.Message = "ok"
	return resp, nil
}

// 获取对象资源 ok
func (s *Server) Get(ctx context.Context, req *pb.GetHttpRequest) (*pb.GetHttpResponse, error) {
	resp := &pb.GetHttpResponse{
		Code: 0,
	}
	err := mongo.Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	// 关闭连接
	defer mongo.Disconnect()

	var result map[string]interface{}
	filter := bson.M{"configId": req.ConfigId}
	projection := bson.M{"_id": 0} //过滤_id字段
	err = mongo.FindOneProjection(config.COLLECTION, filter, &result, projection)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	resultString := string(jsonData)
	resp.Data = resultString
	resp.Message = "ok"
	return resp, nil
}

// 更新对象资源
func (s *Server) Update(ctx context.Context, req *pb.UpdateHttpRequest) (*pb.UpdateHttpResponse, error) {
	resp := &pb.UpdateHttpResponse{
		Code: 0,
	}
	err := mongo.Connect("mongodb://localhost:27017", "testdb")
	if err != nil {
		resp.Message = err.Error()
		return resp, err
	}
	// 关闭连接
	defer mongo.Disconnect()

	logs := make([]map[string]interface{}, 0)
	configId := req.ConfigId
	for _, v := range req.Data {
		data := []byte(v.Data)
		var upData interface{}
		json.Unmarshal(data, &upData)
		path := v.Path
		paths, err := mongo.OperateStr(path)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
		// 构造查询条件
		filter := bson.M{"$and": []bson.M{
			{"configId": configId},
			{path: bson.M{"$exists": true}},
		}}
		// 构造更新条件
		update := bson.M{}

		operateType := paths[0] //操作类型，目前只有两种nodeList basedata
		if v.Action == "d" {    //删除
			if operateType == "nodeList" { //判断path最后一个字符串是否是整型数字，如果是则说明当前操作的是数组
				// // 删除子数组========================================
				// path_p := ""
				// lastDotIndex := strings.LastIndex(path, ".")
				// if lastDotIndex > 0 {
				// 	path_p = path[:lastDotIndex]
				// } else {
				// 	resp.Message = err.Error()
				// 	return resp, err
				// }
				// _, err = mongo.UpdateOne(config.COLLECTION, filter, bson.M{
				// 	"$unset": bson.M{
				// 		path: 1,
				// 	}})
				// if err != nil {
				// 	resp.Message = err.Error()
				// 	return resp, err
				// }
				// update = bson.M{"$pull": bson.M{
				// 	path_p: nil,
				// }}
			} else {
				update = bson.M{
					"$unset": bson.M{
						path: 1,
					},
				}
			}
		} else if v.Action == "e" { //数组调整顺序
			// var slices bson.M
			// // 构造聚合管道
			// pipeline := []bson.M{
			// 	{
			// 		"$match": bson.M{"configId": configId},
			// 	},
			// 	{
			// 		"$project": bson.M{
			// 			path:  1,
			// 			"_id": 0,
			// 		},
			// 	},
			// }
			// err = mongo.Aggregate(config.COLLECTION, pipeline, slices)
			// if err != nil {
			// 	resp.Message = err.Error()
			// 	return resp, err
			// }
		} else {
			if operateType == "nodeList" {
				if len(paths) > 1 {
					nextPath := ""
					for i := 0; i < len(paths); i++ {
						if i == 0 {
							continue
						}
						nextPath += paths[i] + "."
					}
					nextPath = nextPath[:len(nextPath)-1]
					filter = bson.M{"configId": configId, "nodeList.id": v.Id}
					update = bson.M{"$set": bson.M{fmt.Sprintf("%v.$.%v", operateType, nextPath): upData}}
				} else {
					update = bson.M{
						"$set": bson.M{
							path: upData,
						},
					}
				}
			} else {
				update = bson.M{
					"$set": bson.M{
						path: upData,
					},
				}
			}
		}
		// 存在即更新，不存在即插入================================
		actionType := "" //操作类型
		result, err := mongo.UpdateOne(config.COLLECTION, filter, update)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
		// 判断更新结果
		if v.Action == "d" { //删除
			// if result.ModifiedCount == 0 {
			// 	resp.Message = "Not found field!"
			// 	return resp, nil
			// }
			actionType = "delete"
			resp.Message = "Deleted field!"
		} else {
			if result.MatchedCount > 0 {
				actionType = "update"
				resp.Message = "Updated field!"
			} else {
				// 如果不存在，则插入新字段
				_, err = mongo.UpdateOne(config.COLLECTION, bson.M{"configId": configId}, update)
				if err != nil {
					resp.Message = err.Error()
					return resp, err
				}
				actionType = "create"
				resp.Message = "Inserted field!"
			}
		}
		logs = append(logs, map[string]interface{}{
			"actionType": actionType,
			"typeId":     v.TypeId,
			"info":       fmt.Sprintf("%v:%v:%v", path, upData, actionType),
		})
	}
	// 记录操作日志
	mongo.RecordLogs(configId, logs)
	key := fmt.Sprintf("configId_%d", configId)
	recordLogsNum := mongo.IncrLogs(key)
	if recordLogsNum > 0 && (recordLogsNum%config.RECORDLOGSNUM) == 0 {
		//复制一份最新的文档到新的集合中
		var curLog map[string]interface{}
		filter := bson.M{"configId": configId}
		projection := bson.M{"_id": 0} //过滤_id字段
		mongo.FindOneProjection(config.COLLECTION, filter, &curLog, projection)
		curLog["createAt"] = time.Now().Format("2006-01-02 15:04:05")
		_, err = mongo.InsertOne(config.COLLECTION_LOGS, curLog)
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}
	}

	return resp, nil

}

func main() {
	ctx := context.Background() //不带超时时间的ctx，所以不会被取消，除非手动取消
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port)) //监听端口
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC-server服务
	s := grpc.NewServer()
	// 注册gRPC-server服务
	pb.RegisterMongoHttpServer(s, NewServer())
	// 启动gRPC-Server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 创建一个gRPC客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf(":%d", server_port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	m := &runtime.JSONPb{} //定义以哪种数据格式返回给客户端	默认json格式
	// m := &runtime.ProtoMarshaller{} //二进制流格式返回
	// 用于将RESTful API转换成等效的gRPC调用
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, m))

	// 自动注册grpc客户端，并与grpc服务端通信
	err = pb.RegisterMongoHttpHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// 用于创建HTTP请求路由的标准库的方法。它可以用于HTTP服务器端点，用于将不同的HTTP请求映射到不同的处理函数。可以简单理解为这是路由
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", gw_port),
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	// 8088端口提供GRPC-Gateway服务
	log.Println("GRPC-GATEWAY on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, httpServer http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			//grpc-server请求
			grpcServer.ServeHTTP(w, r)
		} else {
			//http-server请求
			httpServer.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
