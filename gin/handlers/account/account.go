package account

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	pb_account "demo/gin/utils/proto/account"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWebRegisterCode(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.UserPhoneRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetWebRegisterCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebSMSLogin(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.UserPhoneCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebSMSLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebLogin(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.LoginRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func WebLogout(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.UserLogOutRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.WebLogout(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetChangePwdCode(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.GetChangePwdCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetChangePwdCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func CheckCode(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.CheckCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.CheckCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func ChangePwdCode(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.ChangePwdCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.ChangePwdCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetEmailCode(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.GetEmailCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.GetEmailCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func CheckEmail(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.CheckEmailRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.CheckEmail(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func BypassAccountChangePWD(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.BypassAccountChangePWDRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.BypassAccountChangePWD(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func AccountInfo(ctx *gin.Context) {
	client := GetClient()
	req := &pb_account.AccountInfoRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil && err != io.EOF {
		ctx.JSON(http.StatusBadRequest, &config.GWResponse{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	res, err := client.AccountInfo(ctx, req)
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
func GetClient() pb_account.AccountHttpClient {
	conn, _ := grpc_client.GetGRPCClient(config.SERVER_NAME1)
	client := pb_account.NewAccountHttpClient(conn)
	return client
}
