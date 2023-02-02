package host

import (
	"net"
	"strings"
)

func GetHostip(prefer string) string {
	addrs, err := net.InterfaceAddrs()
	var ip []string
	var ipret string
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = append(ip, ipnet.IP.String())
				}
			}
		}
	}
	var ipnil string
	for _, ipmem := range ip {
		dataindex := strings.Index(ipmem, prefer)
		if dataindex == 0 {
			ipret = ipmem
		} else {
			ipnil = ipmem
			continue
		}
	}

	if ipret == "" {
		ipret = ipnil
	}
	return ipret
}

func GetHostipList() []string {
	addrs, err := net.InterfaceAddrs()
	var ip []string
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = append(ip, ipnet.IP.String())
				}
			}
		}
	}
	return ip
}
