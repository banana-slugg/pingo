package ip

import (
	"net"
	"testing"

	"github.com/crims1n/pingo/internal/ip"
)

func TestNormalSubnet(t *testing.T) {
	size := 254
	first := "10.0.0.1"
	last := "10.0.0.254"
	test := net.IPNet{
		IP:   net.IPv4(10, 0, 0, 1),
		Mask: net.IPv4Mask(255, 255, 255, 0),
	}
	addrs := ip.GetIPRange(&test)
	if len(addrs) != size {
		t.Errorf("number of addresses should be %d, got %d instead.", size, len(addrs))
	}

	gotFirst := addrs[0]

	if gotFirst.String() != first {
		t.Errorf("the first IP should be %s, got %s instead.", first, gotFirst)
	}

	gotLast := addrs[size-1]

	if gotLast.String() != last {
		t.Errorf("the last IP should be %s, got %s instead.", last, gotLast)
	}

}

func TestBiggerSubnet(t *testing.T) {
	size := 2046
	first := "192.168.8.1"
	last := "192.168.15.254"
	test := net.IPNet{
		IP:   net.IPv4(192, 168, 8, 1),
		Mask: net.IPv4Mask(255, 255, 248, 0),
	}
	addrs := ip.GetIPRange(&test)
	if len(addrs) != size {
		t.Errorf("number of addresses should be %d, got %d instead.", size, len(addrs))
	}

	gotFirst := addrs[0]

	if gotFirst.String() != first {
		t.Errorf("the first IP should be %s, got %s instead.", first, gotFirst)
	}

	gotLast := addrs[size-1]

	if gotLast.String() != last {
		t.Errorf("the last IP should be %s, got %s instead.", last, gotLast)
	}

}

func TestSmallerSubnet(t *testing.T) {
	size := 6
	first := "172.100.1.1"
	last := "172.100.1.6"
	test := net.IPNet{
		IP:   net.IPv4(172, 100, 1, 1),
		Mask: net.IPv4Mask(255, 255, 255, 248),
	}
	addrs := ip.GetIPRange(&test)
	if len(addrs) != size {
		t.Errorf("number of addresses should be %d, got %d instead.", size, len(addrs))
	}

	gotFirst := addrs[0]

	if gotFirst.String() != first {
		t.Errorf("the first IP should be %s, got %s instead.", first, gotFirst)
	}

	gotLast := addrs[size-1]

	if gotLast.String() != last {
		t.Errorf("the last IP should be %s, got %s instead.", last, gotLast)
	}

}
