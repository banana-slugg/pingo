package main

import (
	"fmt"
	"log"

	"github.com/crims1n/pingo/internal/ip"
	"github.com/crims1n/pingo/internal/pinger"
)

func main() {
	ipnet := ip.GetLocalAddr()
	if ipnet == nil {
		log.Fatal("Unable to get local IPv4 address")
	}

	ips := ip.GetIPRange(ipnet)
	res := pinger.PingAll(ips)

	for _, v := range res {
		fmt.Println(v)
	}

}
