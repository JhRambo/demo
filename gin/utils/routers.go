package utils

import (
	"demo/gin/config"
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
		handles += `handler_` + protos[i] + ` "demo/gin/handlers/` + protos[i] + `"
		`
		for _, v := range ms {
			if strings.Contains(config.MSGPACK_URI, strings.ToUpper(v["uri"])) {
				//通用msgpack协议入口，服务端根据uri跳转到对应的服务处理
				uris += `r.` + v["method"] + `("` + v["uri"] + `", handler_msgpack.MsgPackProtocol)
				`
			} else {
				uris += `r.` + v["method"] + `("` + v["uri"] + `", handler_` + protos[i] + `.` + v["rpcName"] + `)
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
}
