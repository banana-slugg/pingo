package ip

import (
	"encoding/binary"
	"log"
	"math"
	"net"
)

func GetIPRange(ip *net.IPNet) []string {
	networkAddr := ip.IP.Mask(ip.Mask)
	addrs := make([]string, 0)
	size := GetSubnetSize(ip)
	buffer := make(net.IP, 4)
	for i := 1; i < int(size); i++ {
		binary.BigEndian.PutUint32(buffer, uint32(i)|binary.BigEndian.Uint32(networkAddr)) // bitwize or
		addrs = append(addrs, buffer.String())
	}
	return addrs
}

func GetSubnetSize(ip *net.IPNet) uint32 {
	return binary.BigEndian.Uint32(ip.Mask) ^ math.MaxUint32 // bitwize xor
}

func GetLocalAddr() *net.IPNet {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		addr, ok := a.(*net.IPNet)               // type assertion to get the actual IP pointer
		isV4 := addr.IP.To4() != nil             // To4() returns nil if the IP is v6
		if ok && !addr.IP.IsLoopback() && isV4 { // if assertion succeeds, IP isn't loopback, and IP isn't a v6 address
			return addr
		}
	}
	return nil
}
