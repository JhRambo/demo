package binary

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_binary "demo/gin/proto/binary"
	"demo/gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	bys, err := utils.GetBinary(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	// 注册gRPC-client客户端
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_binary.NewBinaryHttpClient(conn)
	stream, err := client.UploadFile(ctx)
	req := &pb_binary.BinaryRequest{Data: bys}
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
}
