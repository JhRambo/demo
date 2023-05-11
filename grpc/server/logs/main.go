package main

import (
	pb "demo/grpc/proto/logs"
	"demo/grpc/utils"
	"log"
)

func main() {
	actionType := pb.ActionType_login
	eid := int64(1000)
	uid := int64(2000)
	spaceId := int64(123)
	devId := "222222222222"
	err := utils.Logs(actionType, eid, uid, spaceId, devId)
	if err != nil {
		log.Fatalln(err)
	}
}
