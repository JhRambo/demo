package main

import (
	"fmt"
	"net"
)

func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func main() {
	ip, err := GetLocalIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	apiURL := fmt.Sprintf("http://%s:38401/v2/alarm/feishu/notify1", ip)
	fmt.Println("API URL:", apiURL)
}
