package message

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_message "demo/gin/utils/proto/message"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullProtoMessage1(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.PushInvitedMessageRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.PullProtoMessage1(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func PullProtoMessage2(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.PushInvitedMessageRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.PullProtoMessage2(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// 注册gRPC-client客户端
func GetClient() pb_message.PullServiceClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_message.NewPullServiceClient(conn)
	return client
}
