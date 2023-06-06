package server

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_server "demo/gin/utils/proto/server"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerList(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.GetServerListRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetServerList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetRenderServerList(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.GetRenderServerListRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetRenderServerList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func StartVR(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.StartVRRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.StartVR(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func StartPlanet(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.StartPlanetRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.StartPlanet(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetLogFileList(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.GetLogFileListRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetLogFileList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetInternalList(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.GetInternalListRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetInternalList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func ServerDeviceOut(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.ServerDeviceOutRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.ServerDeviceOut(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func ServersRestartSteam(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.ServersRestartSteamRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.ServersRestartSteam(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func ServersRestartApp(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.ServersRestartAppRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.ServersRestartApp(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func ServersAddr(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.ServersAddrRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.ServersAddr(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func RestartServer(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.RestartServerRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.RestartServer(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetUpdateFiles(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.GetServerUpdateFilesRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetUpdateFiles(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func UpdateServer(ctx *gin.Context) {
	client := GetClient()
	req := &pb_server.UpdateServerRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.UpdateServer(ctx, req)
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
func GetClient() pb_server.ServerHttpClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_server.NewServerHttpClient(conn)
	return client
}
