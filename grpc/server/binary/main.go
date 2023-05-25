package main

import (
	"context"
	pb "demo/grpc/proto/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const server_port = 8081 //grpc-server端口
const gw_port = 8088     //gateway网关端口

type Server struct {
	pb.UnimplementedBinaryHttpServer
}

func NewServer() *Server {
	return &Server{}
}

// gRPC ClientStream 的使用
func (s *Server) UploadFile(stream pb.BinaryHttp_UploadFileServer) error {
	var fileData []byte
	for {
		// 从客户端流中接收数据
		chunk, err := stream.Recv()
		log.Println("err============", err)
		if err == io.EOF { //数据传输结束
			stream.SendAndClose(&empty.Empty{})
			break
		}
		if err != nil {
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}
	log.Println("fileData============", fileData)
	return nil
}

// gw server 监听不同端口
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
	pb.RegisterBinaryHttpServer(s, NewServer())
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

	// m := &runtime.JSONPb{} //定义以哪种数据格式返回给客户端	默认json格式
	m := &runtime.ProtoMarshaller{} //二进制流格式返回
	// 用于将RESTful API转换成等效的gRPC调用
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, m))

	// 自动注册grpc客户端，并与grpc服务端通信
	err = pb.RegisterBinaryHttpHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// 用于创建HTTP请求路由的标准库的方法。它可以用于HTTP服务器端点，用于将不同的HTTP请求映射到不同的处理函数。
	mux := http.NewServeMux()
	mux.Handle("/", middleware(ctx, gwmux, conn))
	// mux.Handle("/", gwmux)

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
		// bys, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }
		// // 手动创建grpc客户端
		// client := pb.NewBinaryHttpClient(conn)
		// stream, err := client.UploadFile(ctx)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }
		// sendBinaryData(stream, bys, w, r)
		// return
		next.ServeHTTP(w, r)
	})
}

// grpc客户端发送二进制流数据
func sendBinaryData(stream pb.BinaryHttp_UploadFileClient, bys []byte, w http.ResponseWriter, r *http.Request) {
	req := &pb.BinaryRequest{Data: bys}
	err := stream.Send(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// 关闭流
	_, err = stream.CloseAndRecv()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
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
