package main

import (
	"bytes"
	"compress/zlib"
	"context"
	pb "demo/grpc/proto/file"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
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
	errorMsg := ""
	stackMsg := ""
	// 创建 zlib 解压缩器
	zlibReader, err := zlib.NewReader(bytes.NewReader(fileData))
	if err != nil {
		return err
	}
	defer zlibReader.Close()

	// 创建一个正则表达式，用于匹配 <ErrorMessage></ErrorMessage> 中的值
	errorMessageRegex := regexp.MustCompile(`<ErrorMessage>(.*?)</ErrorMessage>`)
	stackMessageRegex := regexp.MustCompile(`<CallStack>([\s\S]*?)</CallStack>`)
	uncompressedData, err := ioutil.ReadAll(zlibReader)
	if err != nil {
		return err
	}
	uncompressedString := string(uncompressedData)
	matches := errorMessageRegex.FindAllStringSubmatch(uncompressedString, -1)
	for _, match := range matches {
		if len(match) > 1 {
			errorMsg = match[1]
		}
	}
	matches1 := stackMessageRegex.FindAllStringSubmatch(uncompressedString, -1)
	for _, match := range matches1 {
		if len(match) > 1 {
			if strings.Contains(match[1], "UMyGameInstance::Shutdown()") {
				stackMsg = match[1]
			}
		}
	}
	if errorMsg != "" {
		errorMsg = "<ErrorMessage>" + errorMsg + "</ErrorMessage>"
	}
	if stackMsg != "" {
		stackMsg = "<CallStack>" + stackMsg + "</CallStack>"
	}
	result := errorMsg
	if stackMsg != "" {
		result = result + "\n" + stackMsg
	}
	log.Println("result:", result)
	return nil
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
	pb.RegisterFileHttpServer(s, &Server{})
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
	err = pb.RegisterFileHttpHandler(ctx, gwmux, conn)
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
	log.Println("GRPC-GATEWAY on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}

// 自定义中间件
func middleware(ctx context.Context, next http.Handler, conn *grpc.ClientConn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// bys, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Fatalln("Read failed:", err)
		// }
		// client := pb.NewFileHttpClient(conn)
		// // 初始化客户端流对象
		// stream, err := client.UploadFile(ctx)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }
		// // 发送客户端数据流
		// stream.Send(&pb.FileChunk{
		// 	Data: bys,
		// })
		// // 结束客户端流，等待服务端响应
		// stream.CloseAndRecv()
		// return
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
