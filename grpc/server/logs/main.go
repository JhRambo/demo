package main

import (
	pb "demo/grpc/proto/logs"
	"demo/grpc/utils"
	"log"
)

func main() {
	mp := map[string]interface{}{
		"action_type": pb.ActionType_login,
		"eid":         int64(1000),
		"uid":         int64(2000),
		"space_id":    int64(666),
		"device_id":   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	err := utils.Logs(mp)
	if err != nil {
		log.Fatalln(err)
	}
}
