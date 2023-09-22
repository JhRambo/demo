package main

import (
	"fmt"
	"net"
	"os"
)

func GetHostIP() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ip := net.ParseIP(addr)
		if ip != nil && ip.To4() != nil {
			return ip.String(), nil
		}
	}

	return "", fmt.Errorf("Failed to get host IP")
}

func main() {
	hostIP, err := GetHostIP()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Host IP:", hostIP)
	}
}
