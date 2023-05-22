package main

import (
	"context"
	pb "demo/grpc/proto/file"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const server_port = 8081 //server端口
const gw_port = 8088     //gw网关端口

type Server struct {
	pb.UnimplementedFileHttpServer
}

func NewServer() *Server {
	return &Server{}
}

// gRPC ClientStream 的使用
func (s *Server) UploadFile(stream pb.FileHttp_UploadFileServer) error {
	var fileData []byte
	for {
		// 从客户端流中接收数据
		chunk, err := stream.Recv()
		if err == io.EOF { //数据传输结束
			stream.SendAndClose(&empty.Empty{})
			break
		}
		if err != nil {
			log.Println("err==============", err)
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}
	return nil
}

// gw server 监听不同端口
func main() {
	ctx := context.Background() //不带超时时间的ctx，所以不会被取消，除非手动取消
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册service到server
	pb.RegisterFileHttpServer(s, NewServer())
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

	gwmux := runtime.NewServeMux()
	// 注册HelloHttpHandler
	// ctx2, c := context.WithTimeout(ctx, 10*time.Second) //n秒后超时
	// defer c()
	// var sss grpc.ServerStream
	err = pb.RegisterFileHttpHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	//自定义中间件
	mux := http.NewServeMux()
	// mux.Handle("/", middleware(ctx, gwmux, conn))
	mux.Handle("/", gwmux)

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
		client := pb.NewFileHttpClient(conn)
		// 因为grpc客户端取消了上下文，主要还是因为客户端没有发送数据，导致的================================================
		// ctx, cancel := context.WithTimeout(ctx, 1*time.Nanosecond)
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second) //n秒后超时
		defer cancel()                                         //客户端主动取消
		stream, err := client.UploadFile(ctx)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		sendBinaryData(stream, bys)
		return
		// next.ServeHTTP(w, r)
	})
}

// grpc客户端发送二进制流数据
func sendBinaryData(stream pb.FileHttp_UploadFileClient, bys []byte) error {
	// req := &pb.FileChunk{Data: bys}
	// err := stream.Send(req)
	// if err != nil {
	// 	return err
	// }
	// // 关闭流
	// _, err = stream.CloseAndRecv()
	// if err != nil {
	// 	return err
	// }
	return nil
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
