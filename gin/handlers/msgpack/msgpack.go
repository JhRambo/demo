package msgpack

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_msgpack "demo/gin/proto/msgpack"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MsgPackProtocol(ctx *gin.Context) {
	client := GetClient()
	req := &pb_msgpack.MsgpackHttpRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, _ := client.MsgPackProtocol(ctx, req)
	ctx.JSON(http.StatusOK, res)
}

// 注册gRPC-client客户端
func GetClient() pb_msgpack.MsgpackHttpClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_msgpack.NewMsgpackHttpClient(conn)
	return client
}
