package utils

//初始化路由文件
func InitRouters() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
	handles := ""
	uris := ""
	for i := 0; i < len(protos); i++ {
		protoPath := dir + "/" + protos[i] + ".proto"
		ms := ReadProto(protoPath)
		handles += `handler_` + protos[i] + ` "demo/gin/handlers/` + protos[i] + `"
		`
		for _, v := range ms {
			uris += `r.` + v["method"] + `("` + v["uri"] + `", handler_` + protos[i] + `.` + v["rpcName"] + `)
			`
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
