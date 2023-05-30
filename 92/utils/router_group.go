package utils

/*路由分组
循环生成指定proto包对应的路由分组文件
每个包对应一个文件夹
路径格式：
hello/router.go
binary/router.go
*/
func InitRouterGroup() {
	protos := ScanFiles()
	for i := 0; i < len(protos); i++ {
		content := `
		package ` + protos[i] + `
		import (
			handler_` + protos[i] + " \"demo/gin/handlers/" + protos[i] + `"	
			"github.com/gin-gonic/gin"
		)
		
		// 路由分组
		func InitRouter(r *gin.Engine) {
			g := r.Group("/` + protos[i] + `")
			//这里要通过正则，循环遍历api方法：
			g.POST("/sayhello", handler_` + protos[i] + `.SayHello)
		}	
		`
		filePath := "D:/code/demo/92/routers/" + protos[i] + "/router.go"
		CreateFile(filePath, content)
	}
}
