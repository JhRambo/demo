package routers

import (
	"demo/gin/config"
	pb_binary "demo/gin/proto/binary"
	pb_hello "demo/gin/proto/hello"
	"demo/gin/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var gWResponse = &config.GWResponse{}

func InitRouter(r *gin.Engine, conn *grpc.ClientConn) {
	r.POST("/uploadfile", func(ctx *gin.Context) {
		// 注册gRPC-server客户端
		client := pb_binary.NewBinaryHttpClient(conn)
		stream, err := client.UploadFile(ctx)
		bys, err := ioutil.ReadAll(ctx.Request.Body)
		req := &pb_binary.BinaryRequest{Data: bys}
		err = stream.Send(req)
		if err != nil {
			return
		}
		// 关闭流
		_, err = stream.CloseAndRecv()
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			gWResponse.Code: http.StatusBadRequest,
			gWResponse.Msg:  "ok",
		})
	})
	r.POST("/hello", func(ctx *gin.Context) {
		// 注册gRPC-server客户端
		client := pb_hello.NewHelloHttpClient(conn)
		bys, err := utils.GetBodyBytes(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				gWResponse.Code: http.StatusBadRequest,
				gWResponse.Msg:  err,
			})
			return
		}
		req := &pb_hello.HelloHttpRequest{}
		json.Unmarshal(bys, req)
		res, _ := client.SayHello(ctx, req)
		ctx.JSON(http.StatusOK, res)
	})
}
