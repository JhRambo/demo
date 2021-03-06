package utils

import (
	"regexp"
	"strings"
)

// 生成Handlers
func InitHandlers() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
	for i := 0; i < len(protos); i++ {
		protoPath := dir + "/" + protos[i] + ".proto"
		ms := ReadProto(protoPath)
		fcs := ""
		for _, v := range ms {
			requestParam := strings.TrimSpace(v["request"])
			//binary二进制流特殊处理
			regex := regexp.MustCompile(`stream`)
			if regex.MatchString(v["request"]) {
				re := regexp.MustCompile(`stream\s+(\w+)`)
				match := re.FindStringSubmatch(requestParam)
				if len(match) > 1 {
					requestParam = match[1]
				} else {
					continue
				}
				fcs += `
				func ` + v["rpcName"] + `(ctx *gin.Context) {` + `
					bys, err := utils.GetBinary(ctx)` + `
					if err != nil {
						ctx.JSON(http.StatusOK, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					// 注册gRPC-client客户端
					conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)` + `
					client := pb_` + protos[i] + `.New` + v["serviceName"] + `Client(conn)` + `
					stream, err := client.` + v["rpcName"] + `(ctx)` + `
					req := &pb_` + protos[i] + `.` + requestParam + `{Data: bys}` + `
					err = stream.Send(req)
					if err != nil {
						ctx.JSON(http.StatusOK, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					// 关闭流
					res, err := stream.CloseAndRecv()
					if err != nil {
						ctx.JSON(http.StatusOK, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					ctx.JSON(http.StatusOK, res)
				}`
				content := `
					package ` + protos[i] + `

					import (` + `
						"demo/gin/utils"
						grpc_client "demo/gin/client"` + `
						"demo/gin/config"` + `
						"net/http"` + `
						pb_` + protos[i] + `"demo/gin/proto/` + protos[i] + `"
						"github.com/gin-gonic/gin"` + `
					)
				` + fcs
				filePath := "D:/code/demo/gin/handlers/" + protos[i] + "/" + protos[i] + ".go"
				CreateFile(filePath, content)
			} else {
				fcs += `
				func ` + v["rpcName"] + `(ctx *gin.Context) {` + `
					// 注册gRPC-client客户端 ` + `
					conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)` + `
					client := pb_` + protos[i] + `.New` + v["serviceName"] + `Client(conn)` + `
					req := &pb_` + protos[i] + `.` + requestParam + `{}` + `
					if err := ctx.ShouldBindJSON(req); err != nil {` + `
						ctx.JSON(http.StatusOK, &config.GWResponse{` + `
							Code: -1,` + `
							Message:  err.Error(),` + `
						})` + `
						return` + `
					}` + `
					res, _ := client.` + v["rpcName"] + `(ctx, req)` + `
					ctx.JSON(http.StatusOK, res)` + `
				}
				`
				content := `
					package ` + protos[i] + `

					import (` + `
						grpc_client "demo/gin/client"` + `
						"demo/gin/config"` + `
						"net/http"` + `
						pb_` + protos[i] + `"demo/gin/proto/` + protos[i] + `"
						"github.com/gin-gonic/gin"` + `
					)
				` + fcs
				filePath := "D:/code/demo/gin/handlers/" + protos[i] + "/" + protos[i] + ".go"
				CreateFile(filePath, content)
			}
		}
	}
}
