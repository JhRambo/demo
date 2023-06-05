package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

/*
	初始化proto文件自动生成代码

_grpc.pb.go
.pb.go
*/
func InitProto(protoPath string) {
	protocPath := `D:\MySoft\protoc-23.1-win64\bin\protoc.exe`
	// 获取proto文件列表
	files, err := filepath.Glob(filepath.Join(protoPath, "*.proto"))
	if err != nil {
		fmt.Println("Failed to read proto dir:", err)
	}

	// 遍历proto文件列表
	for _, file := range files {
		// 执行protoc命令生成protobuf代码
		cmdProtoc := exec.Command(protocPath, "-I", protoPath, "--go_out="+filepath.Dir(file), file)
		cmdProtoc.Stdout = os.Stdout
		cmdProtoc.Stderr = os.Stderr
		if err := cmdProtoc.Run(); err != nil {
			fmt.Println("Failed to run protoc command:", err)
		}

		// 执行protoc命令生成gRPC代码
		cmdGRPC := exec.Command(protocPath, "-I", protoPath, "--go-grpc_out="+filepath.Dir(file), file)
		cmdGRPC.Stdout = os.Stdout
		cmdGRPC.Stderr = os.Stderr
		if err := cmdGRPC.Run(); err != nil {
			fmt.Println("Failed to run protoc command:", err)
		}

		fmt.Printf("proto file %s generated successfully\n", file)
	}
}
