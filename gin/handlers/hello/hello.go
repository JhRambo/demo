package hello

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_hello "demo/gin/proto/hello"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SayHello(ctx *gin.Context) {
	// 注册gRPC-client客户端
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_hello.NewHelloHttpClient(conn)
	req := &pb_hello.HelloHttpRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, _ := client.SayHello(ctx, req)
	ctx.JSON(http.StatusOK, res)
}

func SayGoodbye(ctx *gin.Context) {
	// 注册gRPC-client客户端
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_hello.NewHelloHttpClient(conn)
	req := &pb_hello.GoodByeHttpRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, _ := client.SayGoodbye(ctx, req)
	ctx.JSON(http.StatusOK, res)
}
