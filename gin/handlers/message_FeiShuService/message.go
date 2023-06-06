package message

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_message "demo/gin/utils/proto/message"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebFsAlarmPush(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.FsAlarmPushRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebFsAlarmPush(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func FsAlarmPush(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.FsAlarmPushRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.FsAlarmPush(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebFsAlarmRobotSet(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.WebFsAlarmRobotSetRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebFsAlarmRobotSet(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebFsAlarmRobotListGet(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.WebFsAlarmRobotListGetRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebFsAlarmRobotListGet(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebFsAlarmRobotDelete(ctx *gin.Context) {
	client := GetClient()
	req := &pb_message.WebFsAlarmRobotDeleteRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebFsAlarmRobotDelete(ctx, req)
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
func GetClient() pb_message.FeiShuServiceClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_message.NewFeiShuServiceClient(conn)
	return client
}
