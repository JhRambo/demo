package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	pb "demo/grpc/proto/stream"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 定义gRPC服务中的输入和输出结构体
type BytesInput struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

type BytesOutput struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

// 实现gRPC服务
type myServer struct {
	pb.UnimplementedExampleServiceServer
}

func (s *myServer) ProcessBinaryData(stream pb.ExampleService_ProcessBinaryDataServer) error {
	// 从流中读取输入数据
	var input BytesInput
	for {
		if err := stream.RecvMsg(&input); err != nil {
			return err
		}
		// 处理输入数据，并生成输出数据
		output := processBinaryData(input)
		// 将输出数据写回流中
		if err := stream.SendMsg(&output); err != nil {
			return err
		}
	}
}

func processBinaryData(input BytesInput) BytesOutput {
	// 处理输入数据，这里我们简单地将输入数据拼接起来作为输出数据
	return BytesOutput{Data: append([]byte("Processed: "), input.Data...)}
}

// 实现HTTP服务，对应的gRPC服务是myServer
func startHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterExampleServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		return err
	}

	http.ListenAndServe(":8088", mux)
	return nil
}

// 发送HTTP请求
func sendHTTPRequest(data []byte) ([]byte, error) {
	// 将输入数据编码为base64字符串
	inputStr := base64.StdEncoding.EncodeToString(data)
	// 创建HTTP请求对象
	req, err := http.NewRequest("POST", "http://localhost:8088/example/ProcessBinaryData", nil)
	if err != nil {
		return nil, err
	}
	// 添加自定义HTTP头部，包含base64编码的输入数据
	md := metadata.Pairs("x-input-data", inputStr)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 发送HTTP请求
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 从响应中读取并解码输出数据
	outputStr := resp.Header.Get("x-output-data")
	outputData, err := base64.StdEncoding.DecodeString(outputStr)
	if err != nil {
		return nil, err
	}
	return outputData, nil
}

func main() {
	// 启动HTTP服务
	go startHTTPServer()

	// 发送HTTP请求，并处理输出数据
	inputData := []byte("Hello, world!")
	outputData, err := sendHTTPRequest(inputData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output data: %s\n", string(outputData))
}
