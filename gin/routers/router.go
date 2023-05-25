package routers

import (
	pb "demo/grpc/proto/binary"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRouter(r *gin.Engine, conn *grpc.ClientConn) {
	r.POST("/uploadfile", func(c *gin.Context) {
		// 手动创建grpc客户端
		client := pb.NewBinaryHttpClient(conn)
		stream, err := client.UploadFile(c)
		bys, err := ioutil.ReadAll(c.Request.Body)
		req := &pb.BinaryRequest{Data: bys}
		err = stream.Send(req)
		if err != nil {
			return
		}
		// 关闭流
		_, err = stream.CloseAndRecv()
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
}
