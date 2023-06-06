package message

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_message "demo/gin/utils/proto/message"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PushInviteMessage(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.PushInvitedMessageRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.PushInviteMessage(ctx, req)
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
func GetClient() pb_message.PushServiceClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_message.NewPushServiceClient(conn)
	return client
}
