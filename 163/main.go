package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Failed to retrieve network interfaces:", err)
		return
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Failed to retrieve addresses for interface", iface.Name, ":", err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// 判断是否为IPv4地址，可根据需求更改为IPv6
			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.String()
				fmt.Println("Server IP:", ip)
			}
		}
	}
}
