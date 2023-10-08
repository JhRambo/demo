package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "alarm.3dxr.com"

	ipAddresses, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("无法解析域名 %s: %v\n", domain, err)
		return
	}

	for _, ipAddress := range ipAddresses {
		fmt.Printf("%s 的 IP 地址为 %s\n", domain, ipAddress.String())
	}
}
