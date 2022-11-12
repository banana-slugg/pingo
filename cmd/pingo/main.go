package main

import (
	"fmt"
	"log"

	"github.com/crims1n/pingo/internal/ip"
)

func main() {
	ipnet := ip.GetLocalAddr()
	if ipnet == nil {
		log.Fatal("Unable to get local IPv4 address")
	}

	res := ip.GetIPRange(ipnet)
	fmt.Println(res)

}
