package utils

/* 创建router.go文件
该文件初始化路由 */
func InitRouters() {
	protos := ScanFiles()
	filePath := "./routers/router.go"
	routers := make(map[string]string)
	for i := 0; i < len(protos); i++ {
		routers["router_"+protos[i]] = "\"demo/gin/routers/" + protos[i] + "\""
	}
	router_pb := ""
	router_init := ""
	for k, v := range routers {
		router_pb += k + " " + v + "\n"
		router_init += k + ".InitRouter(r)" + "\n"
	}
	content := `
	package routers

	import (
		` + router_pb + `
		"github.com/gin-gonic/gin"
	)

	func InitRouter(r *gin.Engine) {
		` + router_init + `
	}
	`

	CreateFile(filePath, content)
}
