package pinger

import (
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/crims1n/pingo/internal/ip"
	"github.com/go-ping/ping"
)

func PingOnce(ip net.IP) bool {
	isValid := false
	pinger, err := ping.NewPinger(ip.String())
	if err != nil {
		return false
	}
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv > 0 {
			isValid = true
		}
	}
	pinger.Count = 3
	pinger.Timeout = time.Second
	err = pinger.Run()
	if err != nil {
		return false
	}
	return isValid
}

func PingAll(ips []net.IP) []net.IP {
	var wg sync.WaitGroup
	valid := make([]net.IP, 0)
	for _, ip := range ips {
		wg.Add(1)
		go func(ip net.IP) {
			if PingOnce(ip) {
				valid = append(valid, ip)
			}
			defer wg.Done()
		}(ip)
	}

	wg.Wait()
	return ip.SortAddrs(valid)
}
