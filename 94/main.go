package main

import (
	"demo/gin/utils"
)

func main() {
	filename := "D:/code/demo/gin/proto/message.proto"
	protoFile := utils.FilterFile(filename)
	utils.ReadProto(protoFile)
}
