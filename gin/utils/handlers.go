package utils

import (
	"fmt"
)

// 初始化Handlers业务处理文件
func InitHandlers() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
	for i := 0; i < len(protos); i++ {
		fmt.Println("gin generating code from", protos[i]+".proto")
		protoPath := dir + "/" + protos[i] + ".proto"
		protoFile := FilterFile(protoPath)
		ms := ReadProto(protoFile)
		fcs := ""
		for _, v := range ms {
			requestParam := v["request"]
			if v["streamBinary"] == "stream" {
				//stream binary二进制流
				fcs += `
				func ` + v["rpcName"] + `(ctx *gin.Context) {
					bys, err := utils.GetBinary(ctx)
					if err != nil && err != io.EOF {
						ctx.JSON(http.StatusBadRequest, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					client := GetClient()
					stream, err := client.` + v["rpcName"] + `(ctx)
					req := &pb_` + protos[i] + `.` + requestParam + `{` + v["bytesName"] + `: bys}
					err = stream.Send(req)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					// 关闭流
					res, err := stream.CloseAndRecv()
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					ctx.JSON(http.StatusOK, res)
				}`
				content := `
					package ` + protos[i] + `

					import (
						"demo/gin/utils"
						grpc_client "demo/gin/client"
						"demo/gin/config"
						"io"
						"net/http"
						pb_` + protos[i] + `"demo/gin/utils/proto/` + v["packName"] + `"
						"github.com/gin-gonic/gin"
					)
				` + fcs

				content += `
				// 注册gRPC-client客户端
				func GetClient() pb_` + protos[i] + `.` + v["serviceName"] + `Client {
					conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
					client := pb_` + protos[i] + `.New` + v["serviceName"] + `Client(conn)
					return client
				}
				`
				filePath := "D:/code/demo/gin/handlers/" + protos[i] + "/" + protos[i] + ".go"
				CreateFile(filePath, content)
			} else if v["serviceName"] == "MsgpackHttp" {
				// Msgpack协议处理
				fcs += `
				//通用msgpack协议入口，服务端根据uri跳转到对应的服务处理
				func ` + v["rpcName"] + `(ctx *gin.Context) {
					bys, err := utils.GetBinary(ctx)
					if err != nil && err != io.EOF {
						ctx.JSON(http.StatusBadRequest, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					client := GetClient()
					req := &pb_` + protos[i] + `.` + requestParam + `{
						Key: strings.ToUpper(ctx.Request.RequestURI),
						Val: bys,
					}
					res, err := client.` + v["rpcName"] + `(ctx, req)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					ctx.Data(http.StatusOK, "application/x-msgpack", res.Data)
				}
				`
				content := `
					package ` + protos[i] + `

					import (
						"demo/gin/utils"
						grpc_client "demo/gin/client"
						"demo/gin/config"
						"io"
						"net/http"
						pb_` + protos[i] + `"demo/gin/utils/proto/` + v["packName"] + `"
						"github.com/gin-gonic/gin"
						"strings"
					)
				` + fcs

				content += `
				// 注册gRPC-client客户端
				func GetClient() pb_` + protos[i] + `.` + v["serviceName"] + `Client {
					conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
					client := pb_` + protos[i] + `.New` + v["serviceName"] + `Client(conn)
					return client
				}
				`

				filePath := "D:/code/demo/gin/handlers/" + protos[i] + "/" + protos[i] + ".go"
				CreateFile(filePath, content)
			} else {
				//json协议
				fcs += `
				func ` + v["rpcName"] + `(ctx *gin.Context) {
					client := GetClient()
					req := &pb_` + protos[i] + `.` + requestParam + `{}
					if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
						ctx.JSON(http.StatusBadRequest, &config.GWResponse{
							Code: -1,
							Message:  err.Error(),
						})
						return
					}
					res, err := client.` + v["rpcName"] + `(ctx, req)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
							Code:    -1,
							Message: err.Error(),
						})
						return
					}
					ctx.JSON(http.StatusOK, res)
				}
				`
				content := `
					package ` + protos[i] + `

					import (
						grpc_client "demo/gin/client"
						"demo/gin/config"
						"io"
						"net/http"
						pb_` + protos[i] + `"demo/gin/utils/proto/` + v["packName"] + `"
						"github.com/gin-gonic/gin"
					)
				` + fcs

				content += `
				// 注册gRPC-client客户端
				func GetClient() pb_` + protos[i] + `.` + v["serviceName"] + `Client {
					conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
					client := pb_` + protos[i] + `.New` + v["serviceName"] + `Client(conn)
					return client
				}
				`

				filePath := "D:/code/demo/gin/handlers/" + protos[i] + "/" + protos[i] + ".go"
				CreateFile(filePath, content)
			}
		}
	}
}
