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
	"regexp"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = 8081
const pattern_msgpack = "HELLO" //这里追加使用msgpack协议的URI
const CTXINFO = "ctx_info"

var MapNode = map[string]string{}

type Server struct {
	pb.UnimplementedHelloHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloHttpRequest) (*pb.HelloHttpResponse, error) {
	return &pb.HelloHttpResponse{
		Name: req.Name,
		Age:  req.Age,
	}, nil
}

// 使用msgpack协议
func (s *Server) SayBinary(ctx context.Context, req *pb.BinaryRequest) (*pb.BinaryResponse, error) {
	var bys []byte
	if req.Key == "/hello" {
		resp := &pb.HelloHttpResponse{}
		// msgpack.Unmarshal(req.Val, resp)
		resp.Age = 2000
		resp.Name = "玉皇大帝元始天尊女娲娘娘"
		bys, _ = msgpack.Marshal(resp)
	}
	return &pb.BinaryResponse{
		Val: bys,
	}, nil
}

func (s *Server) SayGoodBye(ctx context.Context, req *pb.GoodByeHttpRequest) (*pb.GoodByeHttpResponse, error) {
	d, _ := GetMetaData(ctx)
	obj := &pb.GoodByeHttpRequest{}
	msgpack.Unmarshal(d.N3.Val, obj)
	return &pb.GoodByeHttpResponse{
		Code:  200,
		Token: obj.Token,
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
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 创建一个gRPC客户端连接
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
	mux.Handle("/", middleware(ctx, gwmux, conn))

	gwServer := &http.Server{
		Addr: ":8088",
		// Handler: gwmux,
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	// 8088端口提供gRPC-Gateway服务
	log.Println("gw gRPC-Gateway on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// 定义一个Root结构体内嵌自定义结构体
type NodeRoot struct {
	//这里追加自定义proto===============TODO
	N3 pb.BinaryRequest
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
	nodeRoot := GetNodeRoot()
	strinfo := md[CTXINFO][0]
	json.Unmarshal([]byte(strinfo), &nodeRoot)
	return nodeRoot, nil
}

// 自定义中间件
func middleware(ctx context.Context, next http.Handler, conn *grpc.ClientConn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bys, err := ioutil.ReadAll(r.Body) //获取二进制流
		if err != nil {
			log.Println("Read failed:", err)
		}

		uri := strings.ToUpper(r.RequestURI)

		//1.context请求（方案1）
		nodes := GetNodeRoot()
		nodes.N3.Key = uri
		nodes.N3.Val = bys
		MapNodeData(nodes)

		//2.模拟grpc客户端直接发起grpc请求（方案2）
		re_msgpack := regexp.MustCompile(pattern_msgpack)
		match_msgpack := re_msgpack.MatchString(uri)
		if match_msgpack {
			client := pb.NewHelloHttpClient(conn) //这里改成动态即可 TODO
			byRequest := &pb.BinaryRequest{
				Key: r.RequestURI,
				Val: bys,
			}
			resp, err := client.SayBinary(ctx, byRequest)
			if err != nil {
				log.Fatalln("err:", err.Error())
			}
			v, err := msgpack.Marshal(resp.Val)
			if err != nil {
				log.Fatalln("err:", err.Error())
			}
			w.Write(v)
		}

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
