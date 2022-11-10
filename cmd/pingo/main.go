package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip := GetLocalAddr()
	fmt.Println(ip)
}

func GetLocalAddr() *net.IPNet {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		addr, ok := a.(*net.IPNet) // type assertion to get the actual IP pointer
		isV4 := addr.IP.To4() != nil
		if ok && !addr.IP.IsLoopback() && isV4 { // if assertion succeeds, IP isn't loopback, and IP isn't a v6 address
			return addr
		}
	}
	return nil
}
