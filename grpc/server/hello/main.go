package main

import (
	"context"
	pb "demo/grpc/proto/hello"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = 8081
const CTXINFO = "ctx_info"

var MapNode = map[string]string{}

type Server struct {
	pb.UnimplementedHelloHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloHttpRequest) (*pb.HelloHttpResponse, error) {
	u, _ := GetMetaData(ctx)
	obj := &pb.HelloHttpRequest{}
	msgpack.Unmarshal(u.N3.Val, obj) //二进制流数据转结构体
	return &pb.HelloHttpResponse{
		Name: obj.Name,
		Age:  obj.Age,
	}, nil
}

// gw server 监听不同端口
func main() {
	ctx := context.Background()
	log.Println(fmt.Sprintf("server 监听%d端口...", port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册service到server
	pb.RegisterHelloHttpServer(s, &Server{})
	// 启动gRPC Server
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的gRPC服务器的，客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		ctx,
		":8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux(SetMetaData())
	// 注册HelloHttp
	err = pb.RegisterHelloHttpHandler(ctx, gwmux, conn) //handler
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", middleware(gwmux))

	gwServer := &http.Server{
		Addr: ":8088",
		// Handler: gwmux,
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("gw gRPC-Gateway on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// 定义一个Root结构体内嵌自定义结构体
type NodeRoot struct {
	//这里追加自定义proto===============TODO
	N3 pb.BinaryData
}

func GetNodeRoot() *NodeRoot {
	return &NodeRoot{}
}

// 设置自定义key-value
func MapNodeData(x *NodeRoot) map[string]string {
	u, _ := json.Marshal(x)
	MapNode[CTXINFO] = string(u)
	return MapNode
}

// 自定义metadata
func SetMetaData() runtime.ServeMuxOption {
	return runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		md := metadata.Pairs(CTXINFO, MapNode[CTXINFO])
		return md
	})
}

// 获取context metadata 数据
func GetMetaData(ctx context.Context) (*NodeRoot, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	// log.Println("md================", md)
	nodeRoot := GetNodeRoot()
	strinfo := md[CTXINFO][0]
	json.Unmarshal([]byte(strinfo), &nodeRoot)
	return nodeRoot, nil
}

// 自定义中间件
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body) //获取二进制流
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		nodes := GetNodeRoot()
		nodes.N3.Key = r.RequestURI
		nodes.N3.Val = b
		MapNodeData(nodes)
		next.ServeHTTP(w, r)
	})
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
