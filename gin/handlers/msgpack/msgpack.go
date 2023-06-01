package msgpack

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_msgpack "demo/gin/proto/msgpack"
	"demo/gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
)

func MsgPackProtocol(ctx *gin.Context) {
	bys, err := utils.GetBinary(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	client := GetClient()
	req := &pb_msgpack.MsgpackHttpRequest{
		Key: ctx.Request.RequestURI,
		Val: bys,
	}
	res, err := client.MsgPackProtocol(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	data, err := msgpack.Marshal(res.Val)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.Data(http.StatusOK, "application/x-msgpack", data)
}

// 注册gRPC-client客户端
func GetClient() pb_msgpack.MsgpackHttpClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_msgpack.NewMsgpackHttpClient(conn)
	return client
}