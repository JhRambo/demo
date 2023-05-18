package main

import (
	"context"
	pbhello "demo/grpc/proto/hello"
	pb "demo/grpc/proto/msgpack"
	svrhello "demo/grpc/server/hello"
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

const server_port = 8081                //server端口
const gw_port = 8088                    //gw网关端口
const pattern_msgpack = "HELLO|GOODBYE" //这里追加使用msgpack协议的URI
const CTXINFO = "ctx_info"

var MapNode = map[string]string{}

type Server struct {
	pb.UnimplementedMsgpackHttpServer
}

func NewServer() *Server {
	return &Server{}
}

// 使用msgpack协议
func (s *Server) Binary(ctx context.Context, req *pb.MsgpackRequest) (*pb.MsgpackResponse, error) {
	var bys []byte
	switch req.Key {
	case "/HELLO":
		r := &pbhello.HelloHttpRequest{}
		msgpack.Unmarshal(req.Val, r)
		s := svrhello.NewServer()
		resp, _ := s.SayHello(ctx, r) //跳转到指定的服务去执行
		bys, _ = msgpack.Marshal(resp)
	}
	return &pb.MsgpackResponse{
		Val: bys,
	}, nil
}

// gw server 监听不同端口
func main() {
	ctx := context.Background()
	log.Println("server gRPC-Gateway on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册service到server
	pb.RegisterMsgpackHttpServer(s, &Server{})
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
		fmt.Sprintf(":%d", server_port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux(SetMetaData()) //这里是重点，手动创建的grpc客户端没有这个操作，导致没有往context里塞数据，所以手动创建的grpc客户端也需要设置
	// 注册HelloHttpHandler，这里注册了，下面如果手动注册了grpc客户端，导致同一个方法会被执行两次，需要注意
	err = pb.RegisterMsgpackHttpHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", middleware(ctx, gwmux, conn))

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", gw_port),
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	// 8088端口提供GRPC-Gateway服务
	log.Println("gw gRPC-Gateway on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// 自定义中间件
func middleware(ctx context.Context, next http.Handler, conn *grpc.ClientConn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bys, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Read failed:", err)
		}
		uri := strings.ToUpper(r.RequestURI)
		re_msgpack := regexp.MustCompile(pattern_msgpack)
		match_msgpack := re_msgpack.MatchString(uri)

		if match_msgpack {
			//1.context请求（方案1）
			nodes := GetNodeRoot()
			nodes.N3.Key = uri
			nodes.N3.Val = bys
			MapNodeData(nodes)

			//2.模拟grpc客户端直接发起grpc请求（方案2）
			// 手动注册grpc客户端
			if uri == "/HELLO" {
				client := pb.NewMsgpackHttpClient(conn) //这里改成动态即可 TODO
				byRequest := &pb.MsgpackRequest{
					Key: uri,
					Val: bys,
				}
				ctx := metadata.NewOutgoingContext(ctx, metadata.Pairs(CTXINFO, MapNode[CTXINFO]))
				resp, err := client.Binary(ctx, byRequest)
				if err != nil {
					log.Fatalln("err1:", err.Error())
				}
				v, err := msgpack.Marshal(resp.Val)
				if err != nil {
					log.Fatalln("err2:", err.Error())
				}
				w.Write(v)
				return
			}
		}
		// next.ServeHTTP(w, r)
	})
}

// 定义一个Root结构体内嵌自定义结构体
type NodeRoot struct {
	//这里追加自定义proto===============TODO
	N3 pb.MsgpackRequest
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
