package util

import (
	"net"
)

func CheckIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)

	if ip == nil || ip.To4() == nil {
		return false
	} else {
		return true
	}
}

// 根据网卡接口名获取对应的ipv4的地址
func GetIpByNICName(name string) string {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		panic(err)
	}
	addrs, err := iface.Addrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}
