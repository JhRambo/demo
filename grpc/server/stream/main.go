package main

import (
	"context"
	pb "demo/grpc/proto/stream"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const server_port = 8081 //server端口
const gw_port = 8088     //gw网关端口

type Server struct {
	pb.UnimplementedStreamHttpServer
}

func NewServer() *Server {
	return &Server{}
}

// gRPC ClientStream 的使用
func (s *Server) UploadFile(stream pb.StreamHttp_UploadFileServer) error {
	ctx := context.TODO()
	c := make(chan error, 1)
	var fileData []byte
	for {
		select {
		case <-ctx.Done():
			<-c
			return ctx.Err() // 处理上下文取消信号
		case err := <-c:
			log.Println(err)
		default:
			// 从客户端流中接收数据
			// grpc客户端没有发送数据这里接收不到==================================================
			chunk, err := stream.Recv()
			log.Fatalln(chunk)
			c <- err
			log.Println(grpc.ErrorDesc(err)) //context canceled
			log.Println(err)                 //rpc error: code = Canceled desc = context canceled
			if err == io.EOF {
				stream.SendAndClose(&empty.Empty{})
				break
			}
			if err != nil {
				return err
			}
			fileData = append(fileData, chunk.Data...)
		}
	}
	// return nil
}

// gw server 监听不同端口
func main() {
	ctx := context.Background()
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册service到server
	pb.RegisterStreamHttpServer(s, NewServer())
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

	// m := &runtime.JSONPb{} //定义以哪种数据格式返回给客户端	默认json格式
	// m := &runtime.ProtoMarshaller{} //二进制流格式返回
	// gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, m))
	gwmux := runtime.NewServeMux()

	// 注册HelloHttpHandler
	err = pb.RegisterStreamHttpHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	//自定义中间件
	mux := http.NewServeMux()
	mux.Handle("/", middleware(ctx, gwmux, conn))

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", gw_port),
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	// 8088端口提供GRPC-Gateway服务
	log.Println("GRPC-GATEWAY on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// 自定义中间件
func middleware(ctx context.Context, next http.Handler, conn *grpc.ClientConn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bys, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Read failed:", err)
		}
		sendBinaryData(ctx, conn, bys)
		next.ServeHTTP(w, r)
	})
}

func sendBinaryData(ctx context.Context, conn *grpc.ClientConn, data []byte) error {
	client := pb.NewStreamHttpClient(conn)
	// 初始化客户端流对象
	stream, err := client.UploadFile(ctx)
	// 发送客户端数据流
	stream.Send(&pb.FileChunk{
		Data: data,
	})
	return err
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
