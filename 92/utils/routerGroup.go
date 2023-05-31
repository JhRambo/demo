package utils

import "strings"

/*路由分组
循环生成指定proto包对应的路由分组文件
每个包对应一个文件夹
路径格式：
hello/router.go
binary/router.go
*/
func InitRouterGroup() {
	dir := "D:/code/demo/gin/proto"
	protos := ScanFiles(dir)
	for i := 0; i < len(protos); i++ {
		protoPath := dir + "/" + protos[i] + ".proto"
		ms := ReadProto(protoPath)
		gs := ""
		for _, v := range ms {
			ss := strings.Split(v["uri"], "/")
			uri := ss[len(ss)-1]
			//这里要通过正则，循环遍历api方法：
			gs += "g." + v["method"] + "(\"" + uri + "\", handler_" + protos[i] + "." + v["rpcName"] + ")\n"
		}
		content := `
		package ` + protos[i] + `
		import (
			handler_` + protos[i] + " \"demo/gin/handlers/" + protos[i] + `"	
			"github.com/gin-gonic/gin"
		)
		
		// 路由分组
		func InitRouter(r *gin.Engine) {
			g := r.Group("/` + protos[i] + `/")
			` + gs + `
		}	
		`
		filePath := "D:/code/demo/92/routers/" + protos[i] + "/router.go"
		CreateFile(filePath, content)
	}
}
