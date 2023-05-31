package utils

/* 创建router.go文件
该文件初始化路由 */
func InitRouters() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
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

	// 初始化路由
	func InitRouter(r *gin.Engine) {
		` + router_init + `
	}
	`
	filePath := "D:/code/demo/gin/routers/router.go"
	CreateFile(filePath, content)
}
