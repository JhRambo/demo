package utils

import (
	"demo/gin/config"
	"log"
	"strings"
)

// 初始化路由文件
func InitRouters() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
	handles := ""
	uris := ""
	for i := 0; i < len(protos); i++ {
		protoPath := dir + "/" + protos[i] + ".proto"
		protoFile := FilterFile(protoPath)
		ms := ReadProto(protoFile)

		for _, v := range ms {
			handles += `handler_` + protos[i] + `_` + v["serviceName"] + ` "demo/gin/handlers/` + protos[i] + `_` + v["serviceName"] + `"
			`
			if strings.Contains(config.MSGPACK_URI, strings.ToUpper(v["uri"])) {
				//通用msgpack协议入口，服务端根据uri跳转到对应的服务处理
				uris += `r.` + v["method"] + `("` + v["uri"] + `", handler_msgpack_MsgpackHttp.MsgPackProtocol)
				`
			} else {
				uris += `r.` + v["method"] + `("` + v["uri"] + `", handler_` + protos[i] + `_` + v["serviceName"] + `.` + v["rpcName"] + `)
				`
			}
		}
	}
	content := `
	package routers

	import (
		` + handles + `
		"github.com/gin-gonic/gin"
	)

	// 初始化路由
	func InitRouters(r *gin.Engine) {
		` + uris + `
	}
	`
	filePath := "D:/code/demo/gin/routers/routers.go"
	CreateFile(filePath, content)
	log.Println("routers created and updated successfully.")
}
