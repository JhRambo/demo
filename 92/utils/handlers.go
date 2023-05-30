package utils

// 生成Handlers
func initHandlers() {
	fileName := "handler.go"
	//Handlers
	Handlers := map[string]string{
		"handler_hello": "\"demo/gin/handlers/hello\"",
		// "handler_binary": "\"demo/gin/proto/binary\"",
	}
	handler_pb := ""
	for k, v := range Handlers {
		handler_pb += k + " " + v + "\n"
	}
	content := ``

	CreateFile(fileName, content)
}
