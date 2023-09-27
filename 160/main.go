package main

import (
	consul "demo/159"
	"fmt"
	"time"
)

func main() {
	consulConfig := consul.GetConsulInstance()
	for {
		r := consulConfig.WatchConsulKeyChanges("alarm-robot")
		fmt.Println(r)
		time.Sleep(5 * time.Second)
	}
}
