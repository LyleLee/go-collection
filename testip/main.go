package main

import (
	"fmt"
	"net"
)

func main() {
	ipList := []string{"192.168.1.1/24", "fd04:3e42:4a4e:3381::/64"}
	for i := 0; i < len(ipList); i += 1 {
		ip, ipnet, err := net.ParseCIDR(ipList[i])
		if err != nil {
			fmt.Println("Error", ipList[i], err)
			continue
		}
		fmt.Println(ipList[i], "-> ip:", ip, " net:", ipnet)
	}
}
