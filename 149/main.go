package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	fmt.Println(ip)
}
